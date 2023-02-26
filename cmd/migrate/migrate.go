package migrate

import (
	"os"
	"os/exec"

	"github.com/litsea/logger"
)

var (
	migrationsDir = "db/migrations"
)

func checkAtlas() string {
	cmdPath, err := exec.LookPath("atlas")
	if err != nil {
		logger.Errorf("'atlas' not found for migrate cmd")
		os.Exit(1)
	}

	return cmdPath
}
