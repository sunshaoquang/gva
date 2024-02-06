package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

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
