package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/litsea/logger"
	"github.com/spf13/viper"
)

func SetDefault() {
	viper.SetDefault("telegram.poll-duration", "10s")
	viper.SetDefault("assets.base-url", "http://127.0.0.1:8601")
	viper.SetDefault("assets.host", "0.0.0.0")
	viper.SetDefault("assets.port", 8601)
	viper.SetDefault("db.type", "mysql")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.dbname", "keqing")
	viper.SetDefault("cache.dir", "tmp/cache")
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
		srv := &http.Server{
			Addr:              fmt.Sprintf("%s:%d", host, port),
			ReadHeaderTimeout: 30 * time.Second,
		}
		err := srv.ListenAndServe()
		if err != nil {
			logger.Error("pprof listen failed: ", err.Error())

			return
		}
	}()
}
