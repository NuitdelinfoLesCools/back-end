package server

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler"
	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/middleware"
)

// New .
func New(static string, staticPrefix string) *gin.Engine {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())

	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, object.Fail(fmt.Errorf("Page not found. 404")))
	})

	//setMiddlewares(server)
	constructServer(server)

	return server
}

func constructServer(server *gin.Engine) {
	mdl, err := middleware.JWT()
	if err != nil {
		log.Fatalln(mdl)
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.POST("/user/signup", handler.CreateUser)
	server.POST("/user/auth", mdl.LoginHandler)

	USER := server.Group("/user")
	USER.Use(mdl.MiddlewareFunc())

	POS := USER.Group("/pos")
	POS.GET("/", handler.GetPositions)
	POS.PUT("/", handler.CreatePosition)

	USER.POST("/refresh_token", mdl.RefreshHandler)

	PROTECTED := server.Group("/")
	PROTECTED.Use(mdl.MiddlewareFunc())

	FOOD := PROTECTED.Group("/food")
	FOOD.GET("/", handler.GetFood)
	FOOD.POST("/create", handler.CreateFood)

	ALERT := PROTECTED.Group("/alert")
	ALERT.GET("/", handler.GetAlerts)
	ALERT.POST("/create", handler.CreateAlert)

	MISSION := PROTECTED.Group("/mission")
	MISSION.GET("/", handler.GetMission)
	MISSION.POST("/create", handler.CreateMission)

	EQUIPMENT := PROTECTED.Group("/equipment")
	EQUIPMENT.GET("/", handler.GetEquipement)
	EQUIPMENT.POST("/create", handler.CreateEquipement)

	WEATHER := PROTECTED.Group("/weather")
	WEATHER.GET("/", handler.GetWeather)

}
