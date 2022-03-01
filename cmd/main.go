package main

import (
	"fmt"

	"github.com/seaung/vhub/pkg/config"
	"github.com/seaung/vhub/routers"
)

func main() {
	_ = config.InitConfig()
	port := config.SConfig.App.Port
	fmt.Println("app port : ", port)
	r := routers.InitRouters()
	//r.Run(fmt.Sprintf(":%s", port))
	r.Run(":8000")
}
