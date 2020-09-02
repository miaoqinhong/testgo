package model

import "time"

type UserOrder struct {
	Id int64 `xorm:"pk autoincr" json:"id"`//主键
	SumMoney int64 `xorm:"default 0" json:"sum_money"`//总价格
	Time time.Time `xorm:"DataTime" json:"time"`//时间
	OrderTime uint64 `json:"order_time"`//订单创建时间
	OrderStatusId int64 `xorm:"index" json:"order_status_id"`//订单状态ID
	OrderStatud *OrderStatus `xorm:"-"`
	UserId int64 `xorm:"index" json:"user_id"`//用户编号
	User *User `xorm:"-"`//订单对应的账户，并不进行结构体字段映射
	ShopId  int64  `xorm:"index" json:"shop_id"`//用户购买商品编号
	Shop *Shop `xorm:"-"`//商品结构体，不进行映射
	AddressId int64 `xorm:"index" json:"address_id"`//地址结构体的Id
	Adress *Adderss `xorm:"-"`//地址结构体不进行任何映射
	DelFlag int64 `xorm:"default 0" json:"del_flag"`//删除标志，0代表正常1代表删除
}