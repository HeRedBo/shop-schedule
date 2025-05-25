package po

import (
	"gorm.io/datatypes"
	"time"
)

type ShopUser struct {
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	RealName       string         `json:"real_name"`
	Birthday       int            `json:"birthday"`
	CardId         string         `json:"card_id"`
	Mark           string         `json:"mark"`
	Nickname       string         `json:"nickname"`
	Avatar         string         `json:"avatar"`
	Phone          string         `json:"phone"`
	AddIp          string         `json:"add_ip"`
	LastIp         string         `json:"last_ip"`
	NowMoney       float64        `json:"nowMoney"`
	BrokeragePrice float64        `json:"brokeragePrice"`
	Integral       int            `json:"integral"`
	SignNum        int            `json:"sign_num"`
	Status         int8           `json:"status"`
	Level          int8           `json:"level"`
	SpreadUid      string         `json:"spreadUid"`
	SpreadTime     time.Time      `json:"spread_time"`
	UserType       string         `json:"userType"`
	PayCount       int            `json:"payCount"`
	SpreadCount    int            `json:"spread_count"`
	Address        string         `json:"address"`
	LoginType      string         `json:"login_type"`
	WxProfile      datatypes.JSON `json:"wx_profile"`
	BaseModel
}
