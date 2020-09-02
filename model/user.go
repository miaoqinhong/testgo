package model

import (
	"project_8_26/util"
	"time"
)

/*
用户信息结构体，用于生成用户信息表
*/
type User struct {
	Id int64 `xorm:"pk autoincr" json:"id"`//主键用户ID
	UserName string `xorm:"varchar(12)" json:"user_name"`//用户名称
	RegisterTime time.Time `json:"register_time"`//用户注册时间
	Mobile string `xorm:"varchar(11)" json:"mobile"`//用户手机号
	IsActive int64 `json:"is_active"`//用户是否激活
	Balance string `json:"balance"`//用户账户余额（简单起见用int）
	Avatar string `xorm:"varchar(255)" json:"balance"`//用户头像
	Pwd string `json:"passworld"`//用户密码
	DelFlag int64 `json:"del_flag"`//是否被删除
	CityName string `xorm:"varchar(24)" json:"city_name"`//用户所在城市
	City *City `xorm:"- <- ->"`

}
/*
将数据库数据查询出来的结果进行组装成request请求需要的json字段格式
*/
func (user *User)UserToRespDesc()interface{}  {
	respInfo := map[string]interface{}{
		"id" :          user.Id,
		"user_id":      user.Id,
		"username":     user.UserName,
		"city":         user.CityName,
		"registe_time": util.FormatDatetime(user.RegisterTime),
		"mobile":       user.Mobile,
		"is_active":    user.IsActive,
		"balance":      user.Balance,
		"avatar":       user.Avatar,
	}
	return respInfo
}
