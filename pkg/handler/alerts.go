package handler

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
)

func CreateAlert(c *gin.Context) {
	data := &object.CreateAlert{}

	// Parse request body (json)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	err = store.Agent.CreateAlert(
		data.UserId,
		data.Message,
		data.Severity,
	)

	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(nil))
}

func GetAlerts(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := claims["id"]

	foods, err := store.Agent.GetAlerts(int64(id.(float64)))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(foods))

}
