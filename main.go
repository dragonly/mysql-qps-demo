package main

import (
	"qps_demo/dao"
	"qps_demo/router"
)

func main() {
	dao.InitDB()
	r := router.SetupRouter(dao.DB)
	r.Run()
}
