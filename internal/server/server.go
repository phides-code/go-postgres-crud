package server

import (
	"humdrum/internal/conf"
	"humdrum/internal/database"
	"humdrum/internal/store"
	"log"
	"os"
)

const portKey = "PORT_KEY"

func Start(cfg conf.Config) {
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()
	port, ok := os.LookupEnv(portKey)

	if ok {
		// Start listening and serving requests
		log.Println("*** Server running on port " + port + " ***")
		router.Run(":" + port)
	}
}
