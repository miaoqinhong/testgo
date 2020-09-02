package model


type Shop struct {
	Id int `xorm:"pk autoincr" json:"item_id"`
	Name string `xorm:"varchar(32)" json:"name"`
	Adderss string `xorm:"varchar(128)" json:"adderss"`
	Latitude float32 `json:"latitude"`//经度
	Longitude float32 `json:"longitude"`//纬度
	Description string `xorm:"varchar(255)"`//店铺介绍
	Phone string `json:"phone"`//店铺电话
	PromotionInfo string `json:"promotion_info"`//店铺标语
	FloatDeliveryFee int `json:"float_delivery_fee"` //配送费
	FloatMinimumOrderAmount int `json:"float_minimum_order_amount"`//起送价
	IsPremium bool `json:"is_premium"`//品牌保障
	DeliveryMode bool `json:"delivery_mode"` //蜂鸟专送
	New bool `json:"new"`//是否新店铺
	Bao bool `json:"bao"`//外卖保
	Zhun bool `json:"zhun"`//准时达
	Piao bool `json:"piao"`//是否开发票
	StartTime string `json:"start_time"`//营业开始时间
	EndTime string `json:"end_time"`//营业结束时间
	ImagePath string `json:"image_path"`//店铺头像
	BusinessLicenseImage string `json:"business_license_image"`//店铺营业执照
	CateringServiceLicenseImage string `json:"catering_service_license_image"`//餐饮服务许可证
	Category string `json:"category"`//店铺类型
	Status int `json:"status"`//店铺状态
	RecentOrderNum int `json:"recent_order_num"`//近一个月的销量
	RatingCount int `json:"rating_count"`//评分次数
	Rating	int `json:"rating"`//综合评分
	Dele int `json:"dele"`//是否已经被删除，1代表删除。0代表未删除
	Activities []*Service `xorm:"-"`//商家提供的服务，结构体



}