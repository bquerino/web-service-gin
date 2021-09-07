package main

import (
	"github.com/gin-gonic/gin"
	"mycompany.com/web-service-gin/core/services"
	"mycompany.com/web-service-gin/driven-adapters/repositories"
	"mycompany.com/web-service-gin/driver-adapters/web/http/handlers"
	uidgen "mycompany.com/web-service-gin/utils"
)

func main(){

	albumsRepository := repositories.NewMemKVS()
	albumService := services.New(albumsRepository, uidgen.New())
	albumHandler := handlers.NewHTTPHandler(albumService)

	router := gin.New()
	router.GET("/albums/:id", albumHandler.Get)
	router.POST("/albums", albumHandler.Create)

	router.Run(":8080")
}