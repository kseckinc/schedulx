package main

import (
	"fmt"

	"github.com/galaxy-future/schedulx/api/router"
	"github.com/galaxy-future/schedulx/register/config"
	"github.com/galaxy-future/schedulx/register/config/client"
	"github.com/galaxy-future/schedulx/register/config/log"
)

func main() {
	config.Init("register/conf/config.yml")
	log.Init()
	client.Init()
	r := router.Init()
	err := r.Run(fmt.Sprintf(":%d", config.GlobalConfig.ServerPort))
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
}
