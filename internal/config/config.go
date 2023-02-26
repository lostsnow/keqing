package config

import (
	"fmt"
	"net/http"
	"os"

	"github.com/litsea/logger"
	"github.com/spf13/viper"
)

func SetDefault() {
	viper.SetDefault("telegram.poll-duration", "10s")
	viper.SetDefault("db.type", "mysql")
}

func ReadConfig(cfgFile, configPath string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("app")
		viper.AddConfigPath(configPath)
	}

	return viper.ReadInConfig()
}

func InitLogger() {
	v := viper.Sub("log")

	if err := logger.NewLogger(v, viper.GetString("log_handler")); err != nil {
		fmt.Println("init logger failed: ", err)
		os.Exit(1)
	}

	logger.Info("logger initialized")
}

func InitProfiler(host string, port int) {
	if port <= 0 {
		return
	}

	go func() {
		logger.Infof("pprof listening on %s:%d", host, port)
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
		if err != nil {
			logger.Error("pprof listen failed: ", err.Error())
			return
		}
	}()
}
