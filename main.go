package main

import (
	"app/api"
	"app/dao"
	"app/models"
	"app/pkg/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	dao.Connect()
	models.TableSync()

	g := gin.Default()
	api.InitRouter(g)
	g.Static("/static", "./static")

	err := g.Run(fmt.Sprintf(":%v", config.GetConfig().APP_PORT))
	if err != nil {
		fmt.Println("服务启动失败" + err.Error())
	}

}
