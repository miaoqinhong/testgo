package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"project_8_26/service"
	"project_8_26/util"
)

//每页最大内容
const MaxLimit  = 50
/*
用户控制器结构体：用来实现处理用户模块的接口请求，并返回给客户端
*/

type UserController struct {
	//上下文对象
	Ctx iris.Context
	//user service
	UserService service.UserService

	//session对象
	Session *sessions.Session
}
/*
获取用户总数
请求类型：Get
请求URL：/v1/ueser/count
*/

func (uc *UserController)GetCount() mvc.Result  {
	//用户总数
	total,err := uc.UserService.GetUserTotalCount()
	//如果请求出错
	if err !=nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":util.RECODE_FAIL,
				"count":0,
			},

		}
	}
	//正常返回值
	return mvc.Response{
		Object:map[string]interface{}{
			"status":util.RECODE_OK,
			"count":total,
		},
	}
}

/*
获取用户总数
请求类型：Get
请求url:/v1/users/list
*/
func (uc *UserController)GetList() mvc.Result  {
	offsetStr := uc.Ctx.FormValue("offset")
	limitStr := uc.Ctx.FormValue("limit")
	var offset int
	var limit int
	//判断offset和limit两个变量任意一个都不能为空
	if offsetStr == "" || limitStr==""{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":util.RECODE_FAIL,
				"type":util.RESPMSG_ERROR_USERLIST,
				"message":util.Recode2Text(util.RESPMSG_ERROR_USERLIST),
			},
		}
	}
	//做页面数的限制检查
	if offset <= 0{
		offset= 0
	}
	//做最大的限制
	if limit > MaxLimit{
		limit = MaxLimit
	}
	userList := uc.UserService.GetUserList(offset,limit)

	if len(userList) == 0 {
		return mvc.Response{
			Object:map[string]interface{}{
				"status":util.RECODE_FAIL,
				"type":util.RESPMSG_ERROR_USERLIST,
				"message":util.Recode2Text(util.RESPMSG_ERROR_USERLIST),
			},
		}
	}
	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, user := range userList{
		respList = append(respList,user.UserToRespDesc())
	}
	//返回用户列表
	return mvc.Response{
		Object:&respList,
	}
}