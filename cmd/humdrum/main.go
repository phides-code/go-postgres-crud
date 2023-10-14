package main

import (
	"humdrum/internal/conf"
	"humdrum/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}
