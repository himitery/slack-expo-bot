package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/himitery/slack-expo-bot/src/services/expo"
	"log"
	"os"
)

func Init() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()
	router.Use(jsonMiddleware())

	routerRegistration(router)

	err := router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("[router.Run() Error] %s\n", err)
		return
	}
}

func routerRegistration(router *gin.Engine) {
	router.POST(fmt.Sprintf("/%s", os.Getenv("ENDPOINT")), services.Expo)
}

func jsonMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}
