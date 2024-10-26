package main

import (
	export "export/pkg/router"
	"server/config"
	"server/pkg/env"
	"server/pkg/router"
)

func main() {
	env.Load(config.ExportEnvConfigPath())
	router.ServerRouters(
		"/v1",
		export.ArchiveRouters,
	).Run()
}
