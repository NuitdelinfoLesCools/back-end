package handler

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
)

func CreateMission(c *gin.Context) {
	data := &object.CreateMission{}

	// Parse request body (json)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	//err = store.Agent.CreateMission()

	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(nil))
}

func GetMission(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := claims["id"]

	missions, err := store.Agent.GetMissions(int64(id.(float64)))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(missions))
}
