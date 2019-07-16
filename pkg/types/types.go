package types

type Config struct {
	Uname    string
	//Password string
	//FromEmail string
	//ToEmail  string
	ToUsers  []string
	Duration int64
}

type ItemResult struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Count int    `json:"count"`
	Data  []Item `json:"data"`
}

type Item struct {
	Id               int    `json:"id"`
	ExtendDocument   string `json:"extend_document"`
	AnnounceDocument string `json:"announce_document"`
	GoodsImageUrl    string `json:"goods_image_url"`
	GoodsName        string `json:"goods_name"`
	RefundAmount     string `json:"refund_amount"`
	MinGroupPrice    string `json:"min_group_price"`
	CouponDiscount   string `json:"coupon_discount"`
	StartTime        int64  `json:"start_time"`
	StopTime         int64  `json:"stop_time"`
}

type TokenResult struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}
