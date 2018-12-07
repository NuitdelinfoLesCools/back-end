package plugins

import (
	"math/rand"
	"time"

	owm "github.com/briandowns/openweathermap"
)

var (
	OWAKey string
	lang   string = "EN"
)

type WeatherData struct {
	Temperature float64   `json:"temp"`
	Condition   string    `json:"condition"`
	UVcondition int       `json:"uv"`
	Date        time.Time `json:"date"`
}

func GetLocationInformations(lat, long float64) (*WeatherData, error) {
	w, err := owm.NewCurrent("C", lang, OWAKey)

	if err != nil {
		return nil, err
	}

	err = w.CurrentByCoordinates(&owm.Coordinates{
		Latitude:  lat,
		Longitude: long,
	})

	if err != nil {
		return nil, err
	}

	resp := WeatherData{
		Temperature: w.Main.Temp,
		Condition:   "sun",
		UVcondition: ThisIsNotARandomFunction(1, 6),
	}

	return &resp, nil
}

// ThisIsNotARandomFunction /!\
func ThisIsNotARandomFunction(min, max int) int {
	rand.Seed(time.Now().Unix())
	time.Sleep(50 * time.Millisecond)
	return rand.Intn(max-min) + min
}
