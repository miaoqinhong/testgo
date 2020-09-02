package service

import (
	"fmt"
	"project_8_26/model"
	"xorm.io/xorm"
)

type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)

	//获取管理员总数
	GetAdminCount() (int64, error)
}

func NewAdminService(db *xorm.Engine)AdminService  {
	return &adminService{
		engine:db,

	}
}

type adminService struct {
	engine *xorm.Engine
}

/*
查询管理员总数
*/
func (ac *adminService)GetAdminCount()(int64, error)  {
	count,err := ac.engine.Count(new(model.Admin))
	if err !=nil{
		panic(err.Error())
		return 0,err
	}
	return count,nil

}
/*
通过用户名和密码查询管理员
*/
func (ac *adminService)GetByAdminNameAndPassword(username,password string) (model.Admin,bool) {
	var admin model.Admin
	ac.engine.Where(" admin_name = ? and password = ? ",username,password).Get(&admin)
	fmt.Println("查询后2")
	return admin,admin.AdminId!=0
}