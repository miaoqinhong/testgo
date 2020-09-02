package model

import (
	"fmt"
	"time"
)

type Admin struct {
	AdminId int64 `xorm:"pk autoincr" json:"id"`
	AdminName string `xrom:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xrom:"DateTime" json:"create_time"`
	Status int64 `xorm:"varchar(32)" json:"status"`
	Avatar string `xorm:"varchar(32)" json:"avatar"`
	Pwd string `xorm:"varchar(255)" json:"pwd"`
	CityName string `xorm:"varchar(255)" json:"city_name"`
	CityId int64 `xorm:"index"json:"city_id"`
	City  *City  `xorm:"-<- ->"`//对应所有的城市结构体（基础表结构）
}
/*
从admin数据库试题转换为前端请求的resp的json格式

*/
func (this *Admin)AdminToRespDesc() interface{}  {
	respDesc := map[string]interface{}{
		"user_name":this.AdminName,
		"id":this.AdminId,
		"create_time":this.CreateTime,
		"status":this.Status,
		"avatar":this.Avatar,
		"city":this.CityName,
		"admin":"管理员",
	}
	fmt.Println("获取的数据库信息",respDesc)
	return respDesc
}