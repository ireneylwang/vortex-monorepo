package main

import (
	"server/config"
	"server/pkg/env"
	"server/pkg/router"
)

func main() {
	env.Load(config.PortalEnvConfigPath())
	router.ServerRouters(
		"/v1",
		router.ServerMetricsRouters,
	).Run()
}
