package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"project_8_26/model"
	"xorm.io/core"
	"xorm.io/xorm"
)

func NewMysqlEngine() *xorm.Engine  {
	//数据库引擎
	engine, err := xorm.NewEngine("mysql","root:123@/newproject?charset=utf8")
	if err != nil{
		fmt.Println(err)

	}
	if err :=engine.Ping();err !=nil{
		fmt.Println(err,"数据库连接失败！")
	}


	fmt.Println("数据库连接成功！")
	//自动创建表
	engine.SetMapper(core.GonicMapper{})
	err = engine.Sync2(
		new(model.Admin),
		new(model.City),
		new(model.AdminPermission),
		new(model.User),
		new(model.UserOrder),
		)
	if err !=nil{
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}