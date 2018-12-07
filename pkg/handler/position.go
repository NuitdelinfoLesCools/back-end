package handler

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
)

func CreatePosition(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	data := &object.CreatePosition{}

	// Parse request body (json)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	err = store.Agent.CreatePosition(
		int64(id.(float64)),
		data.Latitude,
		data.Longitude,
		time.Now(),
	)

	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(nil))
}

func GetPositions(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := claims["id"]

	poss, err := store.Agent.GetPositions(int64(id.(float64)))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(poss))
}
