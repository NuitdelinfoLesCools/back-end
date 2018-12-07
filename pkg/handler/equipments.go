package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
)

func CreateEquipement(c *gin.Context) {
	data := &object.CreateEquipement{}

	// Parse request body (json)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	//err = store.Agent.CreateEquipement()

	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(nil))
}

func GetEquipement(c *gin.Context) {
	userId := c.Query("user_id")
	userIdf, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	equipements, err := store.Agent.EquipementStock(int64(userIdf))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(equipements))

}
