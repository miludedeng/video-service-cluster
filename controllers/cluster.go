package controllers

import (
	"encoding/json"
	"fmt"
	"video-service-cluster/cluster"

	"github.com/astaxie/beego"
)

// ClusterController 集群控制器
type ClusterController struct {
	beego.Controller
}

// Index 首页
func (controller *ClusterController) Index() {
	controller.Data["json"] = cluster.C
	controller.ServeJSON(false)
}

// Regist 集群注册方法
func (controller *ClusterController) Regist() {
	var node *cluster.Node
	json.Unmarshal(controller.Ctx.Input.RequestBody, &node)
	fmt.Println(string(controller.Ctx.Input.RequestBody))
	err := cluster.C.Regist(node)
	if err == nil {
		controller.Data["json"] = map[string]interface{}{
			"message": "success",
		}
	} else {
		controller.Data["json"] = map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
		}

	}
	controller.ServeJSON(false)
}

// HealthCheck 健康检测
func (controller *ClusterController) HealthCheck() {
}

// HealthCheck 心跳检测
func (controller *ClusterController) HeartCheck() {
	controller.Data["json"] = map[string]interface{}{
		"status": "success",
	}
	controller.ServeJSON(false)
}
