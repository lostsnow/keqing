package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"github.com/lostsnow/keqing/pkg/entity"
)

var (
	DB *Client
)

type Client struct {
	Client *entity.Client
}

func (m *Client) Init() error {
	client, err := entity.Open(viper.GetString("db.type"), GetDSN())
	if err != nil {
		return fmt.Errorf("open database failed: %w", err)
	}
	defer client.Close()
	//ctx := context.Background()
	// run the auto migration tool
	//err = client.Schema.Create(
	//	ctx,
	//	migrate.WithForeignKeys(false),
	//	migrate.WithDropIndex(true),
	//	migrate.WithDropColumn(true),
	//)
	//if err != nil {
	//	return fmt.Errorf("failed creating schema resources: %w", err)
	//}

	m.Client = client
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
