package cmd

import (
	"fmt"
	"net/http"

	"github.com/litsea/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lostsnow/keqing/data"
)

var assetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "serve assets",
	Run: func(cmd *cobra.Command, args []string) {
		http.Handle("/", http.FileServer(http.FS(data.Asset)))
		addr := fmt.Sprintf("%s:%d", viper.GetString("app.assets.host"), viper.GetInt("app.assets.port"))
		logger.Infof("serve assets to %s", addr)
		err := http.ListenAndServe(addr, nil)
		logger.Errorf("serve assets failed: %s", err)
	},
}
