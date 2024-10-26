package main

import (
	datamagnet "datamagnet/pkg/router"
	iam "iam/pkg/router"
	"server/config"
	"server/pkg/env"
	"server/pkg/router"
)

func main() {
	env.Load(config.OpenApiEnvConfigPath())
	router.ServerRouters(
		"/v1",
		iam.AccessRouters,
		datamagnet.IntegrationRouters,
	).Run()
}
