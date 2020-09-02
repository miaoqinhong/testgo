package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	config2 "project_8_26/config"
	"project_8_26/controller"
	"project_8_26/datasource"
	"project_8_26/model"
	"project_8_26/service"
	"time"

	//"github.com/kataras/iris/v12/mvc"
	//"github.com/kataras/iris/v12/sessions"
	//"newProject/controller"
	//"newProject/service"
	//"oneProject/Text4/ProjectWeb/datasource"
	//"time"
)

func main()  {
	app := newApp()

	//应用App设置
	configtion(app)


	//路由设置
	mvcHandle(app)

	config := config2.InitConfig()
	fmt.Println("配置文件在main",config)
	addr := ":"+config.Port
	app.Run(
		iris.Addr(addr),//在端口8080进行监听
		iris.WithoutServerError(iris.ErrServerClosed),//无服务错误提示
		iris.WithOptimizations,//对json数据序列化更快的配置
		)
	//app.Run(iris.Addr(":8080"))






}
/*
mvc架构模式处理
*/
func mvcHandle(app *iris.Application)  {
	sessManager := sessions.New(sessions.Config{
		Cookie:"sessioncookie",//设置加盐
		Expires:24 *time.Hour,//时长
	})
	engine := datasource.NewMysqlEngine()
	dataEmpty,err := engine.IsTableEmpty(new(model.Admin))
	if err != nil{
		panic(err.Error())
	}
	if dataEmpty{
		fmt.Println("人员表为空")
	}else {
		fmt.Println("人员表不为空")
	}
	//管理员用户模块
	adminService := service.NewAdminService(engine)
	admin := mvc.New(app.Party("/admin"))
	fmt.Println("admin接口",admin)
	admin.Register(
		adminService,
		sessManager.Start,
		)
	admin.Handle(new(controller.AdminController))
	//用户功能模块
	userService := service.NewAdminService(engine)
	user := mvc.New(app.Party("/v1/users"))
	user.Register(
		userService,
		sessManager.Start,
		)
	user.Handle(new(controller.UserController))

	//统计功能模块
	statisService := service.NewAdminService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{date}/"))
	statis.Register(
		statisService,
		sessManager.Start,
		)
	statis.Handle(new(controller.StatisController))

}

func newApp()*iris.Application  {
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	//注册静态资源
	app.HandleDir("/static","./static")
	app.HandleDir("/manage/static","./static")
	app.HandleDir("/img","./static/img")

	//注册视图文件
	app.RegisterView(iris.HTML("./static",".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})
	return app
}


func configtion(app *iris.Application)  {
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset:"UTF-8",
	}))
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg":iris.StatusNotFound,
			"msg":"not found",
			"data":iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg":iris.StatusInternalServerError,
			"msg":"interal error",
			"data":iris.Map{},
		})
	})
}
