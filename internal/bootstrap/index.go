package bootstrap

import (
	"context"
	"learning-keycloak/internal/config"
	"learning-keycloak/internal/helper"
	"learning-keycloak/internal/infrastructure/mysql"
	"learning-keycloak/internal/infrastructure/redis"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	_ = godotenv.Load()

	globalContext := context.Background()

	config.LoadConfig()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:        []string{"*"},
		AllowMethods:        []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:        []string{"Origin", "Accept", "Content-Type", "Authorization", "X-Requested-With", "X-Request-ID"},
		ExposeHeaders:       []string{"Content-Length", "Content-Type"},
		AllowPrivateNetwork: true,
	}))

	redis.NewConnectToRedis(globalContext)
	mysql.NewConnectToMySQL()

	app.GET("/", func(ctx *gin.Context) {
		helper.NewResponseGlobal(ctx, 200, "The Application Is Running Well", nil, nil)
	})

	port := config.NewAppConfig()

	app.Run(":" + port)
}
