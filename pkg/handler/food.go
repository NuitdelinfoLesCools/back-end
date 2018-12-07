package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
)

func CreateFood(c *gin.Context) {
	data := &object.CreateFood{}

	// Parse request body (json)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	err = store.Agent.CreateFood(
		data.UserId,
		data.Name,
		data.Stock,
		data.Quantity,
		data.Unit,
	)

	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(nil))
}

func GetFood(c *gin.Context) {
	userId := c.Query("user_id")
	userIdf, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	foods, err := store.Agent.FoodStock(int64(userIdf))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(foods))

}
