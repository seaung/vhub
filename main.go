package main

import (
	"fmt"

	"github.com/seaung/vhub/pkg/config"
	"github.com/seaung/vhub/routers"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("cause error %+v\n", err)
	}
	//port := config.SConfig.App.Port
	//fmt.Println("app port : ", port)
	r := routers.InitRouters()
	r.Run(fmt.Sprintf(":%s", config.SConfig.App.Port))
}
