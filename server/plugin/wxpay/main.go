package wxpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxpay/router"
	"github.com/gin-gonic/gin"
)

type WxpayPlugin struct {
}

func CreateWxpayPlug(
	MchID string,
	AppId string,
	MchCertificateSerialNumber string,
	MchAPIv3Key string,
	PemPath string,
	NotifyUrl string,
) *WxpayPlugin {
	global.GlobalConfig.MchID = MchID
	global.GlobalConfig.MchCertificateSerialNumber = MchCertificateSerialNumber
	global.GlobalConfig.MchAPIv3Key = MchAPIv3Key
	global.GlobalConfig.PemPath = PemPath
	global.GlobalConfig.AppID = AppId
	global.GlobalConfig.NotifyUrl = NotifyUrl
	return &WxpayPlugin{}
}

func (*WxpayPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitWxpayRouter(group)
}

func (*WxpayPlugin) RouterPath() string {
	return "wxpay"
}
