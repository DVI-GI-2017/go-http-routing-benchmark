package main

import (
	"log"

	"time"

	_ "github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/routes"
)

func init() {
	pool.InitWorkers()
}

func main() {
	config, err := configs.FromFile("config.json")
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	db.InitDB(config.Mongo)

	router, err := mux.NewRouter("/api/v1")
	if err != nil {
		log.Fatalf("can not create router: %v", err)
	}
	router.AddWrappers(mux.Logger, mux.Timeout(5*time.Second))

	routes.SetupRoutes(router)

	router.PrintRoutes()

	port := config.Server.GetPort()

	log.Printf("Server started on port %s...\n", port)

	log.Fatal(router.ListenAndServe(port))
}
