package migrate

import (
	"context"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/litsea/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lostsnow/keqing/internal/db"
	"github.com/lostsnow/keqing/pkg/entity/migrate"
)

var (
	migrationName string
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate migrations",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		// Create a local migration directory able to understand Atlas migration file format for replay.
		dir, err := atlas.NewLocalDir(migrationsDir)
		if err != nil {
			logger.Errorf("failed creating atlas migration directory: %v", err)
			os.Exit(1)
		}
		// Migrate diff options.
		opts := []schema.MigrateOption{
			schema.WithDir(dir),                            // provide migration directory
			schema.WithMigrationMode(schema.ModeInspect),   // provide migration mode
			schema.WithDialect(viper.GetString("db.type")), // Ent dialect to use
			schema.WithDropColumn(true),
			schema.WithDropIndex(true),
			schema.WithForeignKeys(false),
			schema.WithFormatter(atlas.DefaultFormatter),
		}

		// Generate migrations using Atlas support for DB (note the Ent dialect option passed above).
		err = migrate.NamedDiff(ctx, db.GetMigrateDSN(), migrationName, opts...)
		if err != nil {
			logger.Fatalf("failed generating migration file: %v", err)
			os.Exit(1)
		}
		logger.Infof("generate migration %s successfully", migrationName)
	},
}

func init() {
	GenerateCmd.Flags().StringVar(&migrationName, "name", "", "migration name")
	GenerateCmd.MarkFlagRequired("name")
}
