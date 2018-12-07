package handler

import (
	"strconv"

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
	userId := c.Query("user_id")
	userIdf, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	missions, err := store.Agent.GetMissions(int64(userIdf))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(missions))

}
