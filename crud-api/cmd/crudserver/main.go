package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/artemxgod/11-go-projects/crudserver/internal/app/crudserver"
)

var configpath = "configs/crudserver.toml"

func main() {
	cfg := crudserver.NewConfig()
	toml.DecodeFile(configpath, cfg)
	fmt.Printf("Server lanched on localhost%s\n", cfg.BindAddr)

	if err := crudserver.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
