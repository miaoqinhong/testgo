package model
type City struct {
	CityId int64 `xorm:"pk autoincr" json:"id"`
	CityName string `xorm:"varchar(12)" json:"name"`
	PinYin string `xorm:"varchar(32)" json:"pin_yin"`
	Longitude float32 `xorm:"default 0" json:"longitude"`
	Latitude string `xorm:"default 0" json:"latitude"`
	AreaCode string `xorm:"varchar(6)" json:"area_code"`
	Abbr string `xorm:"varchar(12)" json:"abbr"`
}