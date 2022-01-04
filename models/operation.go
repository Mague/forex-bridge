package models

type Operation struct {
	Id          int     `gorm:"primaryKey;autoIncrement;type:int;unique" json:"id"`
	AccountId   int     `json:"TR_ID" binding:"required"`
	OrderNumber int     `gorm:"index:acountid,unique";json:"POS_NUM" binding:"required"`
	Symbol      string  `gorm:"size:50" json:"SYMBOL" binding:"required"`
	OrderType   int     `json:"ORD_TYPE" binding:"required"`
	Price       float64 `json:"PRICE" binding:"required"`
	StopLoss    float64 `json:"SL" binding:"required"`
	TakeProfit  float64 `json:"TP" binding:"required"`
}
