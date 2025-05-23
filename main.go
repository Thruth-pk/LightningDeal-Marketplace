package main

import (
	"LightningDeal_Marketplace/backend/web/controllers"
	"LightningDeal_Marketplace/common"
	"LightningDeal_Marketplace/repositories"
	"LightningDeal_Marketplace/services"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	//1.创建iris实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	template := iris.HTML("./backend/web/views", ".html").
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//4.设置模板目录
	app.HandleDir("/assets", "./backend/web/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//5.注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))
	//6.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations)
}
