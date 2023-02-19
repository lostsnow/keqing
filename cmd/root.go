package cmd

import (
	"fmt"
	"github.com/lostsnow/keqing/internal/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	profilerHost string
	profilerPort int
)

var rootCmd = &cobra.Command{
	Use:   "keqing",
	Short: "Keqing bot",
	RunE: func(*cobra.Command, []string) error {
		return fmt.Errorf("invalid command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./configs/app.yml)")
	rootCmd.PersistentFlags().StringVar(&profilerHost, "profiler-host", "127.0.0.1", "profiler host")
	rootCmd.PersistentFlags().IntVar(&profilerPort, "profiler-port", 0, "profiler port")
}

func initConfig() {
	if err := config.ReadConfig(cfgFile, "./configs"); err == nil {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
		viper.WatchConfig()
		config.SetDefault()
		config.InitLogger()
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	config.InitProfiler(profilerHost, profilerPort)
}
