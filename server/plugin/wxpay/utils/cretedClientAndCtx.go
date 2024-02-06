package utils

import (
	"context"
	global2 "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/global"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
)

func CreateClientAndCtx() (error, context.Context, *core.Client) {
	var (
		mchID                      string = global.GlobalConfig.MchID                      // 商户号
		mchCertificateSerialNumber string = global.GlobalConfig.MchCertificateSerialNumber // 商户证书序列号
		mchAPIv3Key                string = global.GlobalConfig.MchAPIv3Key                // 商户APIv3密钥
	)
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(global.GlobalConfig.PemPath)
	if err != nil {
		log.Fatal("load merchant private key error")
	}
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		e := err.Error()
		global2.GVA_LOG.Error("new wechat pay client err:" + e)
		return err, ctx, client
	}
	return err, ctx, client
}
