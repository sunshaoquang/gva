package service

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	wx_global "github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/utils"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	utils2 "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"gorm.io/gorm"
	"log"
	"math"
	"strconv"
	"time"
)

type WxpayService struct{}

func (e *WxpayService) GetPayCode(order model.Order) (err error, CodeUrl string, orderID string) {
	err, ctx, client := utils.CreateClientAndCtx()
	if err != nil {
		return err, "", ""
	}
	err, codeUrl, orderID := payOne(ctx, client, order)
	return err, codeUrl, orderID
}

func payOne(ctx context.Context, client *core.Client, order model.Order) (err error, codeUrl string, orderID string) {
	// 下单用户ID
	// 下单单号（总ID 用6位（100000）开始记录）
	// 下单产品名(Description)
	// 下单价格(Total)分
	// 得到prepay_id，以及调起支付所需的参数和签名
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		//先创建一个空的订单
		//然后 订单一定要记录好订单上面关联的产品ID 包含产品的介绍 名字 金额等

		no := "123456"                // 六位数的预下单ID
		b := int(math.Floor(1 * 100)) // 产品价格  分
		Name := "GVA插件"

		ctx = context.WithValue(ctx, "no", no)
		svc := native.NativeApiService{Client: client}
		resp, _, err := svc.Prepay(ctx,
			native.PrepayRequest{
				Appid:         core.String(wx_global.GlobalConfig.AppID),
				Mchid:         core.String(wx_global.GlobalConfig.MchID),
				Description:   core.String(Name), // 产品名字
				OutTradeNo:    core.String(no),   // 预下单的那个ID 必须大于6位 建议 ID自增从 100000开始
				TimeExpire:    core.Time(time.Now()),
				Attach:        core.String("自定义数据说明"),                        //可以不写
				NotifyUrl:     core.String(wx_global.GlobalConfig.NotifyUrl), // 回调url
				SupportFapiao: core.Bool(false),                              // 开不开发票
				Amount: &native.Amount{
					Currency: core.String("CNY"),
					//Total:    core.Int64(commodity.Place), //商品价格(分)
					Total: core.Int64(int64(b)),
				},
			},
		)
		if err != nil {
			log.Println(err)
			return err
		}
		orderID = no
		codeUrl = *resp.CodeUrl
		return err
	})
	return err, codeUrl, orderID
}

// 关闭订单  用于紧急关闭订单操作 传入订单ID即可
func (e *WxpayService) ClosePayCode(order model.Order) (err error) {
	err, ctx, client := utils.CreateClientAndCtx()
	if err != nil {
		return err
	}
	err = closeOrder(ctx, client, order)
	return err
}

func closeOrder(ctx context.Context, client *core.Client, order model.Order) error {
	svc := native.NativeApiService{Client: client}
	no := strconv.Itoa(int(order.ID))
	result, err := svc.CloseOrder(ctx,
		native.CloseOrderRequest{
			OutTradeNo: core.String(no),
			Mchid:      core.String(wx_global.GlobalConfig.MchID),
		},
	)
	if err != nil {
		// 处理错误
		log.Printf("call CloseOrder err:%s", err)
		return err
	} else {
		// 处理返回结果
		log.Printf("status=%d", result.Response.StatusCode)
		return nil
	}
}

// 查询订单 传入订单ID即可
func (e *WxpayService) GetOrderById(orderID string) (error, model.Order) {
	err, ctx, client := utils.CreateClientAndCtx()
	if err != nil {
		return err, model.Order{}
	}
	return queryOrderByOutTradeNo(ctx, client, orderID)
}

func queryOrderByOutTradeNo(ctx context.Context, client *core.Client, orderID string) (err error, order model.Order) {
	svc := native.NativeApiService{Client: client}
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		native.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String(orderID),
			Mchid:      core.String(wx_global.GlobalConfig.MchID),
		},
	)
	if err != nil {
		// 错误处理
		log.Printf("call QueryOrderByOutTradeNo err:%s", err)
		return err, order
	} else {

		// TradeState
		//SUCCESS：支付成功
		//REFUND：转入退款
		//NOTPAY：未支付
		//CLOSED：已关闭
		//REVOKED：已撤销（仅付款码支付会返回）
		//USERPAYING：用户支付中（仅付款码支付会返回）
		//PAYERROR：支付失败（仅付款码支付会返回）

		// 处理错误
		// 可以根据订单返回的结果做一些业务逻辑

		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
		return err, order
	}
}

func (e *WxpayService) PayAction(pay model.PayAction) error {
	p, err := utils2.DecryptAES256GCM(wx_global.GlobalConfig.MchAPIv3Key, pay.Resource.AssociatedData, pay.Resource.Nonce, pay.Resource.Ciphertext)
	if err != nil {
		global.GVA_LOG.Info(p)
		global.GVA_LOG.Info(err.Error())
		return err
	}
	var payOrder model.PayOrder
	if err != nil {
		global.GVA_LOG.Info(p)
		global.GVA_LOG.Info(err.Error())
		return err
	}
	err = json.Unmarshal([]byte(p), &payOrder)
	// payOrder 为回调信息 请自行根据回调信息做业务逻辑
	return err
}
