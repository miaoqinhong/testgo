package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"project_8_26/model"
	"project_8_26/service"
	"project_8_26/util"
)

//管理员控制器

type AdminController struct {
	//iris 框架自动为每个请求都绑定上下文对象


	Ctx iris.Context
	//admin功能实体

	Service service.AdminService


	//session对象
	Session *sessions.Session
}
const(
	ADMINTABLENAME ="admin"
	ADMIN          ="admin"
)
//管理员登录
type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//管理员退出功能,
//请求类型GET，
//请求URL：admin/singout
func (ac *AdminController)GetSingout()mvc.Result  {
	//删除session，下次需要重新登录
	ac.Session.Delete(ADMIN);
	return mvc.Response{
		Object: map[string]interface{}{
			"status":util.RECODE_OK,
			"success" : util.Recode2Text(util.RESPMSG_SIGNOUT),
		},
	}
}
/*
处理获取管理员总数的路由请求
请求类型：Get
请求Url：admin/count

*/
func (ac *AdminController) GetCount() mvc.Result  {
	count,err := ac.Service.GetAdminCount()
	if err !=nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":util.RECODE_FAIL,
				"message":util.Recode2Text(util.RESPMSG_ERRORADMINCOUNT),
				"count":0,
			},
		}
	}
	return mvc.Response{
		Object:map[string]interface{}{
			"status":util.RECODE_OK,
			"count":count,
		},
	}
}
/*
获取管理员信息接口
请求类型Get
请求Url：/admin/info
*/
func (ac *AdminController)GetInfo() mvc.Result  {
	//从session中获取信息
	userByte := ac.Session.Get(ADMIN)
	//session为空
	if userByte == nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":util.RECODE_UNLOGIN,
				"type":util.EEROR_UNLOGIN,
				"message":util.Recode2Text(util.EEROR_UNLOGIN),
			},
		}
	}
	//解析数据到admin数据结构
	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte),&admin)

	//解析失败
	if err != nil{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":util.RECODE_UNLOGIN,
				"type":util.EEROR_UNLOGIN,
				"message":util.Recode2Text(util.EEROR_UNLOGIN),
			},
		}
	}
	//解析成功
	return mvc.Response{
		Object:map[string]interface{}{
			"status":util.RECODE_OK,
			"data":admin.AdminToRespDesc(),
		},
	}
}

/*
管理员登录功能
接口：/admin/login
*/

func (ac *AdminController)PostLogin(context iris.Context)mvc.Result {
	iris.New().Logger().Info("admin login")
	var adminLogin AdminLogin
	ac .Ctx.ReadJSON(&adminLogin)
	fmt.Println("登录接口！！！")
	fmt.Println(adminLogin.UserName)
	fmt.Println(adminLogin.Password)
	//数据参数检验
	if adminLogin.UserName == "" || adminLogin.Password == ""{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":"0",
				"success":"登录失败",
				"message":"用户名或密码为空，请重新填写后尝试登录",

			},
		}
	}
	//根据用户名，密码到数据库查询对应的管理信息
	fmt.Println("管理员登录查询数据库前")
	admin,exist := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName,adminLogin.Password)
	fmt.Println("管理员登录查询数据库后")
	if !exist{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":"0",
				"successs":"登录失败",
				"message":"用户名或密码错误，请重新登录",
			},
		}
	}
	//管理员存在，设置session
	userByte,_ := json.Marshal(admin)
	ac.Session.Set(ADMIN,userByte)
	return mvc.Response{
		Object:map[string]interface{}{
			"status":"1",
			"successs":"登录成功",
			"message":"管理员登录成功",
		},
	}

}