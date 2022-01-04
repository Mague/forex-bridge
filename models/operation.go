package models

type Operation struct {
	Id          uint    `gorm:"primaryKey;autoIncrement;type:int;unique" json:"id"`
	AccountId   uint    `json:"TR_ID" binding:"required"`
	OrderNumber float32 `json:"POS_NUM" binding:"required"`
	Symbol      string  `gorm:"size:50" json:"SYMBOL" binding:"required"`
	OrderType   uint    `json:"ORD_TYPE" binding:"required"`
	Price       float32 `json:"PRICE" binding:"required"`
	StopLoss    float32 `json:"SL" binding:"required"`
	TakeProfit  float32 `json:"TP" binding:"required"`
}
