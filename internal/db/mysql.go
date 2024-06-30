package db

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/spf13/viper"

	"github.com/lostsnow/keqing/pkg/entity"
)

var DB *Client

type Client struct {
	Client *entity.Client
}

func (m *Client) Init() error {
	drv, err := sql.Open(viper.GetString("db.type"), GetDSN())
	if err != nil {
		return fmt.Errorf("open database failed: %w", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	m.Client = entity.NewClient(entity.Driver(drv))
	DB = m

	return nil
}

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		viper.GetString("db.username"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetInt("db.port"),
		viper.GetString("db.dbname"),
	)
}

func GetMigrateDSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		viper.GetString("db.type"),
		viper.GetString("db.username"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetInt("db.port"),
		viper.GetString("db.dbname"),
	)
}
