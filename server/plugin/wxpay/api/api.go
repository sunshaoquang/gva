package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/model"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WxpayApi struct{}

// @Tags Wxpay
// @Summary 获取微信支付二维码和ID
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/getPayCode[post]
func (p *WxpayApi) GetPayCode(c *gin.Context) {
	var order model.Order
	order.CustomerID = utils.GetUserID(c)
	c.ShouldBindJSON(&order)
	if err, codeUrl, codeId := service.ServiceGroupApp.GetPayCode(order); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("获取付款码失败:"+err.Error(), c)
	} else {
		response.OkWithData(gin.H{
			"codeUrl": codeUrl,
			"codeId":  codeId,
		}, c)
	}
}

// @Tags Wxpay
// @Summary 获取支付结果
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/getOrderById[get]
func (p *WxpayApi) GetOrderById(c *gin.Context) {
	id := c.Query("orderID")
	if err, data := service.ServiceGroupApp.GetOrderById(id); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("获取付款码失败:"+err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}
}

// @Tags Wxpay
// @Summary 获取支付结果
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxpay/payAction[post]
func (p *WxpayApi) PayAction(c *gin.Context) {
	var pay model.PayAction
	c.ShouldBindJSON(&pay)
	if err := service.ServiceGroupApp.PayAction(pay); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		c.JSON(200, gin.H{
			"code":    "FAIL",
			"message": "失败",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "SUCCESS",
			"message": "接收成功",
		})
	}
}
