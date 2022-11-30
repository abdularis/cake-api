package main

import (
	"cake-api/config"
	"cake-api/core"
	"cake-api/handler"
	"cake-api/repository"
	"cake-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Get()
	db, err := utils.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	cakeRepo := repository.NewCakeRepository(db)
	cakeUseCase := core.NewCakeUseCase(cakeRepo)
	cakeHandler := handler.NewCakeHandler(cakeUseCase)

	route := gin.Default()
	api := route.Group("/api/v1")
	api.GET("/cakes", cakeHandler.GetListCakes)
	api.GET("/cakes/:cakeID", cakeHandler.GetCakeDetail)
	api.POST("/cakes", cakeHandler.CreateNewCake)
	api.PATCH("/cakes/:cakeID", cakeHandler.UpdateCake)
	api.DELETE("/cakes/:cakeID", cakeHandler.DeleteCakeByID)

	addr := fmt.Sprintf(":%d", cfg.HostPort)
	log.Info("Running http server at %s\n", addr)
	if err := route.Run(addr); err != nil {
		log.Fatal(err)
	}
}
