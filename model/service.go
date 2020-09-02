package model
type Service struct {
	Id int `xorm:"varchar(32)" json:"id"`
	Name string `xorm:"varchar(32)" json:"name"`
	IconColor string `xorm:"json:"icon_color""`
	Description string `xorm:"varchar(255)" json:"description"`
	Shop []*Shop `xorm:"-"`
}