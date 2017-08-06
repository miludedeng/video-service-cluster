package main

import (
	_ "video-service-cluster/router"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
