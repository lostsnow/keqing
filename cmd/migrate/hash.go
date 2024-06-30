package migrate

import (
	"os"
	"os/exec"

	"github.com/litsea/logger"
	"github.com/spf13/cobra"
)

var HashCmd = &cobra.Command{
	Use:   "hash",
	Short: "update migrations hash",
	Run: func(_ *cobra.Command, _ []string) {
		cmdPath := checkAtlas()
		c := exec.Command(cmdPath, "migrate", "hash",
			"--dir", "file://"+migrationsDir,
		)

		out, err := c.CombinedOutput()
		if err != nil {
			logger.Errorf("update migrations hash failed: %s", out)
			os.Exit(1)
		}
		logger.Infof("update migrations hash successfully")
	},
}
