package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/NuitdelinfoLesCools/back-end/pkg/plugins"
	"github.com/NuitdelinfoLesCools/back-end/pkg/server"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/db"
)

func main() {
	// Config path flag
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "config file path")
	flag.Parse()

	// Loading configuration
	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("can't find config file %s: %v", configPath, err)
	}

	// Init database
	dbEngine, err := store.New(db.ConnStr{
		Host:             config.Database.Host,
		Port:             config.Database.Port,
		User:             config.Database.User,
		Password:         config.Database.Password,
		Database:         config.Database.Name,
		MaxIdleConns:     100,
		MaxOpenConns:     100,
		MaxLifetimeConns: 0.2,
	})

	if err != nil {
		log.Fatalln(err)
	}
	store.Agent = dbEngine

	// Sync store model: create tables if doesn't exist
	err = store.Agent.Sync()
	if err != nil {
		log.Println(err)
	}

	// Try to create fake data
	{
		if u, err := store.Agent.GetUser("test"); err == nil || u == nil {
			store.Agent.CreateUser("test@test.test", "test", "test")
		}
		if missions, err := store.Agent.GetMissions(1); err == nil && len(*missions) == 0 {
			store.Agent.CreateMission(1, "Mission amazon", "brazil amazon mission", time.Now().Add(time.Hour*240), false)
			store.Agent.CreateMission(1, "Impossible Mission", "impossible mission", time.Now().Add(time.Hour*24000), false)
		}

		if positions, err := store.Agent.GetPositions(1); err == nil && len(*positions) == 0 {
			store.Agent.CreatePosition(1, 1.9, 1.66666, time.Now().Add(0*time.Hour))
			store.Agent.CreatePosition(1, 1.6, 1.65555, time.Now().Add(1*time.Hour))
			store.Agent.CreatePosition(1, 1.5, 1.69999, time.Now().Add(2*time.Hour))
			store.Agent.CreatePosition(1, 1.5, 1.89999, time.Now().Add(3*time.Hour))
			store.Agent.CreatePosition(1, 1.7, 1.90000, time.Now().Add(4*time.Hour))
			store.Agent.CreatePosition(1, 1.6, 1.92000, time.Now().Add(5*time.Hour))
		}

		if food, err := store.Agent.FoodStock(1); err == nil && len(*food) == 0 {
			store.Agent.CreateFood(1, "Water", true, 5, "Litre")
			store.Agent.CreateFood(1, "Garlic", true, 1, "Kilo")
			store.Agent.CreateFood(1, "Potatoes", true, 1, "Kilo")
			store.Agent.CreateFood(1, "Tomatoes", true, 1, "Kilo")
			store.Agent.CreateFood(1, "Bread", true, 1, "Kilo")
		}

		if equips, err := store.Agent.EquipementStock(1); err == nil && len(*equips) == 0 {
			store.Agent.CreateEquipement(1, "screwdriver", true)
			store.Agent.CreateEquipement(1, "backpack", true)
			store.Agent.CreateEquipement(1, "gps", true)
			store.Agent.CreateEquipement(1, "camera", true)
		}

		if alerts, err := store.Agent.GetAlerts(1); err == nil && len(*alerts) == 0 {
			store.Agent.CreateAlert(1, "Hello alert", 0)
			store.Agent.CreateAlert(1, "Danger Alert", 0)
		}

	}

	// Set API keys
	plugins.OWAKey = config.ApiKeys.OpenWeatherApi
	//plugins.GetLocationInformations(48.5839200, 7.7455300),
	//fmt.Println(config.ApiKeys.OpenWeatherApi)

	// Create server
	svr := server.New(config.StaticFiles, config.StaticURL)
	svr.Run(":" + strconv.Itoa(config.Port))
}
