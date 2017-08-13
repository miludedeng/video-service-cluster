package router

import (
	"video-service-cluster/controllers"

	"github.com/astaxie/beego"
)

// InitRouter 路由初始化
func init() {
	beego.Router("/cluster/", &controllers.ClusterController{}, "get:Index")
	beego.Router("/cluster/regist", &controllers.ClusterController{}, "post:Regist")
	beego.Router("/cluster/health", &controllers.ClusterController{}, "get:HealthCheck")
	beego.Router("/cluster/heart", &controllers.ClusterController{}, "get:HeartCheck")
}
