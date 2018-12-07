package middleware

import (
	"log"
	"time"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func JWT() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"id":       v.Id,
					"username": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Id:       int64(claims["id"].(float64)),
				Username: claims["username"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals object.AuthUser
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			if ok, user, err := store.Agent.AuthentificateUser(username, password); err == nil && ok {
				return &model.User{
					Id:       user.Id,
					Username: user.Username,
					Email:    user.Email,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return authMiddleware, nil
}
