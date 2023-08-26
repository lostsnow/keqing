package migrate

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/litsea/logger"
	"github.com/spf13/cobra"

	"github.com/lostsnow/keqing/internal/db"
)

var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cmdPath := checkAtlas()
		c := exec.Command(cmdPath, "migrate", "apply",
			"--dir", fmt.Sprintf("file://%s", migrationsDir),
			"--url", db.GetMigrateDSN(),
		)

		out, err := c.CombinedOutput()
		if err != nil {
			logger.Errorf("apply migrations failed: %s", out)
			os.Exit(1)
		}
		logger.Infof("apply migrations successfully: %s", out)
	},
}
