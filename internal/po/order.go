package po

import "time"

type StoreOrder struct {
	OrderId                string    `json:"order_id"`
	ExtendOrderId          string    `json:"extend_order_id"`
	Uid                    int64     `json:"uid"`
	RealName               string    `json:"real_name"`
	UserPhone              string    `json:"user_phone"`
	UserAddress            string    `json:"user_address"`
	CartId                 string    `json:"cart_id"`
	FreightPrice           float64   `json:"freight_price"`
	TotalNum               int       `json:"total_num"`
	TotalPrice             float64   `json:"total_price"`
	TotalPostage           float64   `json:"total_postage"`
	PayPrice               float64   `json:"pay_price"`
	PayPostage             float64   `json:"pay_postage"`
	DeductionPrice         float64   `json:"deduction_price"`
	CouponId               int64     `json:"coupon_id"`
	CouponPrice            float64   `json:"coupon_price"`
	Paid                   int       `json:"paid"`
	PayTime                time.Time `json:"pay_time"`
	PayType                string    `json:"pay_type"`
	Status                 int       `json:"status"`
	RefundStatus           int       `json:"refund_status"`
	RefundReasonWapImg     string    `json:"refund_reason_wap_img"`
	RefundReasonWapExplain string    `json:"refund_reason_wap_explain"`
	RefundReasonTime       time.Time `json:"refund_reason_time"`
	RefundReasonWap        string    `json:"refund_reason_wap"`
	RefundReason           string    `json:"refund_reason"`
	RefundPrice            float64   `json:"refund_price"`
	DeliverySn             string    `json:"delivery_sn"`
	DeliveryName           string    `json:"delivery_name"`
	DeliveryType           string    `json:"delivery_type"`
	DeliveryId             string    `json:"delivery_id"`
	GainIntegral           int       `json:"gain_integral"`
	UseIntegral            int       `json:"use_integral"`
	PayIntegral            int       `json:"pay_integral"`
	BackIntegral           int       `json:"back_integral"`
	Mark                   string    `json:"mark"`
	Unique                 string    `json:"unique"`
	Remark                 string    `json:"remark"`
	CombinationId          int64     `json:"combination_id"`
	PinkId                 int64     `json:"pink_id"`
	Cost                   float64   `json:"cost"`
	SeckillId              int64     `json:"seckill_id"`
	BargainId              int64     `json:"bargain_id"`
	VerifyCode             string    `json:"verify_code"`
	StoreId                int64     `json:"store_id"`
	ShippingType           int       `json:"shipping_type"`
	UserDto                *ShopUser `json:"user_dto" gorm:"foreignKey:Uid;"`
	CartInfo               []Cart    `json:"cart_info" gorm:"-"`
	OrderStatus            int       `json:"_status" gorm:"-"`
	OrderStatusName        string    `json:"status_name" gorm:"-"`
	PayTypeName            string    `json:"pay_type_name" gorm:"-"`
	BaseModel
}

type BaseModel struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	IsDel      bool      `json:"is_del" gorm:"softDelete:flag"`
}

type OrderIndex struct {
	OrderId       string    `json:"order_id"`
	OrderIdSuffix string    `json:"order_id_suffix"`
	Names         []string  `json:"names"`
	ProductIds    []int64   `json:"product_ids"`
	Uid           int64     `json:"uid"`
	PayTime       time.Time `json:"pay_time"`
	PayType       string    `json:"pay_type"`
	RefundStatus  int       `json:"refund_status"` // '0 未退款 1 申请中 2 已退款'
	ShippingType  int       `json:"shipping_type"` //配送方式
	OrderStatus   int       `json:"order_status"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}
