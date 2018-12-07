package handler

import (
	"strconv"

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
	userId := c.Query("user_id")
	userIdf, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	foods, err := store.Agent.GetAlerts(int64(userIdf))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(foods))

}
