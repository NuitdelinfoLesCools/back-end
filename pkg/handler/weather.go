package handler

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/NuitdelinfoLesCools/back-end/pkg/handler/object"
	"github.com/NuitdelinfoLesCools/back-end/pkg/plugins"
)

func GetWeather(c *gin.Context) {
	//FIXME /!\
	//claims := jwt.ExtractClaims(c)
	//id := claims["id"]

	/*equipements, err := store.Agent.EquipementStock(int64(id.(float64)))
	if err != nil {
		c.JSON(200, object.Fail(err))
		return
	}*/

	w := []plugins.WeatherData{
		plugins.WeatherData{
			Temperature: float64(plugins.ThisIsNotARandomFunction(30, 40)),
			Condition:   "sun",
			UVcondition: plugins.ThisIsNotARandomFunction(5, 9),
			Date:        time.Now(),
		},
		plugins.WeatherData{
			Temperature: float64(plugins.ThisIsNotARandomFunction(30, 40)),
			Condition:   "sun",
			UVcondition: plugins.ThisIsNotARandomFunction(5, 9),
			Date:        time.Now().Add(time.Hour * 24),
		},
		plugins.WeatherData{
			Temperature: float64(plugins.ThisIsNotARandomFunction(30, 40)),
			Condition:   "sun",
			UVcondition: plugins.ThisIsNotARandomFunction(5, 9),
			Date:        time.Now().Add(time.Hour * 48),
		},
	}

	c.JSON(200, object.Success(w))
}
