package api

import (
	"github.com/access-module/api/controller"
	"github.com/access-module/api/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel) // NEW
}

func Server() {
	router := gin.Default()

	//Universal middleware (loggers for example)
	// router.Use(middleware.Authentication())

	// Routes not requiring authentication
	router.POST("/user", controller.CreateUser)

	// Apply the Authentication middleware only to routes that require authorization
	router.GET("/user", middleware.Authentication(), controller.ListUser)
	router.Run(":8080")

}
