package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/NuitdelinfoLesCools/back-end/pkg/server"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store"
	"github.com/NuitdelinfoLesCools/back-end/pkg/store/db"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "config file path")
	flag.Parse()

	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("can't find config file %s: %v", configPath, err)
	}

	// Init database module
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

	// Create server
	svr := server.New(config.StaticFiles, config.StaticURL)
	svr.Run(":" + strconv.Itoa(config.Port))
}
