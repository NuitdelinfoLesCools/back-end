package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
)

// CreateUser page handler
func CreateUser(c *gin.Context) {
	signupData := &object.CreateUser{}

	// Parse request body (json)
	err := c.ShouldBindJSON(&signupData)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	err = store.Agent.CreateUser(
		signupData.Email,
		signupData.Username,
		signupData.Password,
	)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	c.JSON(200, object.Success(signupData))
}

func AuthUser(c *gin.Context) {
	authData := &object.CreateUser{}

	err := c.ShouldBindJSON(&authData)
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}

	ok, _, err := store.Agent.AuthentificateUser(
		authData.Username,
		authData.Password,
	)

	if err != nil {
		c.JSON(200, object.Fail(fmt.Errorf("error while looking for user %v: %v", authData.Username, err)))
		return
	}

	if !ok {
		c.JSON(200, object.Fail(fmt.Errorf("username or password incorrect")))
		return
	}

	c.JSON(200, object.Success(ok))
}
