package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/artemxgod/11-go-projects/web-server/internal/app/webserver"
)

var configPath = "./configs/webserver.toml"

func main() {
	cfg := webserver.NewConfig()
	toml.DecodeFile(configPath, &cfg)

	fmt.Printf("Server is launched on 127.0.0.1%s\n", cfg.BindAddr)
	webserver.Start(cfg)
}