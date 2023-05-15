package app

import (
	"github.com/shop/internal/config"
	"github.com/shop/internal/transport"
)

func CreateApp() {
	config.ParseDatabaseEnviroment()
	transport.ServerApp()
}
