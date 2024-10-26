package config

import (
	"path/filepath"
	"runtime"
)

func EnvConfigPath() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(f), "./.env")
}

func PortalEnvConfigPath() string {
	return EnvConfigPath() + ".portal"
}

func OpenApiEnvConfigPath() string {
	return EnvConfigPath() + ".openapi"
}

func ExportEnvConfigPath() string {
	return EnvConfigPath() + ".export"
}
