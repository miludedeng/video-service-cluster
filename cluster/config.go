package cluster

import "github.com/astaxie/beego"

// ClusterName 集群名称
var ClusterName string
var Master string
var NodeAddr string
var NodeCheckTryTimes int

func InitConfig() {
	ClusterName = beego.AppConfig.String("cluster-name")
	Master = beego.AppConfig.String("cluster-master")
	NodeAddr = beego.AppConfig.String("node-addr")
	NodeCheckTryTimes, _ := beego.AppConfig.Int("node-check-try-times")
	if NodeCheckTryTimes == 0 {
		NodeCheckTryTimes = 3
	}
}
