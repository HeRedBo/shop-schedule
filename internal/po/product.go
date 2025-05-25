package po

type Product struct {
	Id        int64   `json:"id"`
	StoreName string  `json:"store_name"`
	CateId    int     `json:"cate_id"`
	Price     float64 `json:"price"`
	VipPrice  float64 `json:"vip_price"`
	OtPrice   float64 `json:"ot_price"`
}
