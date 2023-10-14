package server

import (
	"humdrum/internal/conf"
	"humdrum/internal/database"
	"humdrum/internal/store"
)

func Start(cfg conf.Config) {
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
