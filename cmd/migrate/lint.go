package migrate

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/litsea/logger"
	"github.com/spf13/cobra"
)

var lintLatest int

var LintCmd = &cobra.Command{
	Use:   "lint",
	Short: "lint migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cmdPath := checkAtlas()
		if lintLatest < 1 {
			lintLatest = 1
		}
		c := exec.Command(cmdPath, "migrate", "lint",
			"--dev-url", "docker://mysql/8/test",
			"--dir", fmt.Sprintf("file://%s", migrationsDir),
			"--latest", strconv.Itoa(lintLatest),
		)

		out, err := c.CombinedOutput()
		if err != nil {
			logger.Errorf("lint migrations failed: %s", out)
			os.Exit(1)
		}
		logger.Infof("lint migrations successfully")
	},
}

func init() {
	LintCmd.Flags().IntVar(&lintLatest, "latest", 1, "lint latest")
}
