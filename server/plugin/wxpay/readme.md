## GVA 邮件发送功能插件
#### 开发者：Mr.奇淼

### 使用步骤

#### 1. 前往GVA主程序下的initialize/plugin.go 添加本插件
    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
	PluginInit(PublicGroup, wxpay.CreateWxpayPlug(
		global.GVA_CONFIG.Wxpay.MchID,
		global.GVA_CONFIG.Wxpay.AppID,
		global.GVA_CONFIG.Wxpay.MchCertificateSerialNumber,
		global.GVA_CONFIG.Wxpay.MchAPIv3Key,
		global.GVA_CONFIG.Wxpay.PemPath,
		global.GVA_CONFIG.Wxpay.NotifyUrl,
	))

    同样也可以再传入时写死
    PluginInit(PublicGroup, wxpay.CreateWxpayPlug(
        "MchID",
        "AppID",
        "MchCertificateSerialNumber",
        "MchAPIv3Key",
        "PemPath",
        "NotifyUrl",
    ))

### 2. 配置说明

#### 2-1 全局配置结构体说明

    type Wxpay struct {
        MchID                      string // 商户ID
        AppID                      string // 绑定小程序的APPID
        NotifyUrl                  string // 支付回调域名
        MchCertificateSerialNumber string // 商户证书序列号
        MchAPIv3Key                string // 商户APIv3密钥
        PemPath                    string // 证书文件所在地址
    }

#### 2-2 入参结构说明
```go
//推荐的订单结构体

type Order struct {
    global.GVA_MODEL
    CustomerID  uint // 下单用户ID
    CommodityID uint // 商品ID用于查询商品价格
    // Transaction
    TradeState     string // 订单状态
    TradeStateDesc string // 订单状态详情
    TradeType      string // 订单状态类型
    TransactionId  string // 订单ID

	// 下面为微信的附加信息
	Mchid       string    `json:"mchid"`
	Appid       string    `json:"appid"`
	OutTradeNo  string    `json:"out_trade_no"`
	BankType    string    `json:"bank_type"`
	Attach      string    `json:"attach"`
	SuccessTime time.Time `json:"success_time"`

	Openid string `json:"openid"`

	Total         int    `json:"total"`
	PayerTotal    int    `json:"payer_total"`
	Currency      string `json:"currency"`
	PayerCurrency string `json:"payer_currency"`
}

func (*Order) TableName() string {
    return "shop_order"
}

// 微信回调的结构体
type PayAction struct {
    ID           string    `json:"id"`
    CreateTime   time.Time `json:"create_time"`
    ResourceType string    `json:"resource_type"`
    EventType    string    `json:"event_type"`
    Summary      string    `json:"summary"`
    Resource     struct {
    OriginalType   string `json:"original_type"`
    Algorithm      string `json:"algorithm"`
    Ciphertext     string `json:"ciphertext"`
    AssociatedData string `json:"associated_data"`
    Nonce          string `json:"nonce"`
    } `json:"resource"`
}

// 微信回调解密的结构体
type PayOrder struct {
    Mchid          string    `json:"mchid"`
    Appid          string    `json:"appid"`
    OutTradeNo     string    `json:"out_trade_no"`
    TransactionId  string    `json:"transaction_id"`
    TradeType      string    `json:"trade_type"`
    TradeState     string    `json:"trade_state"`
    TradeStateDesc string    `json:"trade_state_desc"`
    BankType       string    `json:"bank_type"`
    Attach         string    `json:"attach"`
    SuccessTime    time.Time `json:"success_time"`
    Payer          struct {
    Openid string `json:"openid"`
    } `json:"payer"`
    Amount struct {
    Total         int    `json:"total"`
    PayerTotal    int    `json:"payer_total"`
    Currency      string `json:"currency"`
    PayerCurrency string `json:"payer_currency"`
    } `json:"amount"`
}
    
    func (*PayOrder) TableName() string {
    return "shop_order"
    }
```


### 3. 方法API

    utils.CreateClientAndCtx() (error, context.Context, *core.Client)
    创建携带权限的请求微信支付的请求客户端

### 4. 可直接调用的接口

    预下单： /wxpay/getPayCode [post] 需要根据业务设计传入商品信息
    
    获取订单状态  /wxpay/getOrderById[get] {orderID:0}

    支付回调 /wxpay/payAction[post] 供微信使用

 
## 前端（涉及业务 代码并不统一 请按照以下业务逻辑设计）

    一、发起api请求，携带商品ID（千万别带金额等，不要把金钱等敏感信息控制在前端）给后端，调用后端/wxpay/getPayCode接口，创建预下单 获取订单号和二维码
    
    二、将二维码渲染到下面提供的两个二维码工具 生成二维码的同时，创建轮询 用第一步获得的ID持续调用/wxpay/getOrderById[get] 参数：{orderID:第一步给的ID}
    （前端可以用setInterval开启定时器
        const flag(定时器唯一标识) = setInterval(()=>{要做的事},毫秒)
    ）
    
    三、当第二步获取订单信息中 status状态成为 SUCCESS时，发放商品等。

react
```
npm i --save qrcode.react
```

```react
import QRCode  from 'qrcode.react'
// 使用组件即可
<QRCode
value={ order.url }          // 需要生成二维码图片的url地址
size={ 150 }         // 二维码图片大小
fgColor="#000000"  // 二维码图片背景色
/>
```

vue
```
npm i --save qrcode.vue
```

```
// js 部分

import QrcodeVue from 'qrcode.vue'

// template 部分
 <qrcode-vue 
 :value="value"     // 需要生成二维码图片的url地址
 :size="size"  // 二维码图片大小
 />
```

