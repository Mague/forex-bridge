package payloads

type Operation []struct {
	AccountId   uint    `json:"TR_ID"`
	OrderNumber float32 `json:"POS_NUM"`
	Symbol      string  `json:"SYMBOL"`
	OrderType   uint    `json:"ORD_TYPE"`
	Price       float32 `json:"PRICE"`
	StopLoss    float32 `json:"SL"`
	TakeProfit  float32 `json:"TP"`
}
