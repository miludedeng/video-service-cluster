package router

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// InitRouter 路由初始化
func init() {
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.ContentType("application/json")
		ctx.Output.Body([]byte(`
{
  "name": "hello"
}
      `))
	})

	beego.Router("/regist", &controllers.ClusterController{}, "put:Regist")
}
