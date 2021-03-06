package handler

import (
	jwt "github.com/appleboy/gin-jwt"
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
	claims := jwt.ExtractClaims(c)
	id := claims["id"]

	equipements, err := store.Agent.EquipementStock(int64(id.(float64)))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(equipements))

}
