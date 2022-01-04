package payloads

type Operation []struct {
	AccountId   int     `json:"TR_ID"`
	OrderNumber int     `json:"POS_NUM"`
	Symbol      string  `json:"SYMBOL"`
	OrderType   int     `json:"ORD_TYPE"`
	Price       float64 `json:"PRICE"`
	StopLoss    float64 `json:"SL"`
	TakeProfit  float64 `json:"TP"`
}
