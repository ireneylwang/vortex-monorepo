package test

import (
	"os"
	"server/config"
	"server/pkg/env"
	"testing"
)

func TestMain(m *testing.M) {
	env.Load(config.OpenApiEnvConfigPath())
	code := m.Run()
	os.Exit(code)
}
