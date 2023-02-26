package migrate

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/litsea/logger"
	"github.com/spf13/cobra"
)

var HashCmd = &cobra.Command{
	Use:   "hash",
	Short: "update migrations hash",
	Run: func(cmd *cobra.Command, args []string) {
		cmdPath := checkAtlas()
		c := exec.Command(cmdPath, "migrate", "hash",
			"--dir", fmt.Sprintf("file://%s", migrationsDir),
		)

		out, err := c.CombinedOutput()
		if err != nil {
			logger.Errorf("update migrations hash failed: %s", out)
			os.Exit(1)
		}
		logger.Infof("update migrations hash successfully")
	},
}
