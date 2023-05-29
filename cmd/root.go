package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/abdil1234/test-golang/config"
	"github.com/abdil1234/test-golang/internal/app/commons"
	"github.com/abdil1234/test-golang/internal/app/commons/appcontext"
	"github.com/abdil1234/test-golang/internal/app/commons/loggers"
	"github.com/abdil1234/test-golang/internal/app/domain/repositories"
	"github.com/abdil1234/test-golang/internal/app/server/http"
	"github.com/abdil1234/test-golang/internal/app/usecases"
	"github.com/abdil1234/test-golang/internal/app/usecases/game"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var logger = loggers.GetLogger(context.Background())

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my_project",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.GetConfig()
	app := appcontext.NewAppContext(cfg)

	db, err := initDatabase(app, cfg)
	if err != nil {
		logger.Fatalf("Failed to init database | %v", err)
		return
	}

	opt := commons.Option{
		AppCtx:   app,
		Config:   cfg,
		Database: db,
	}

	repo := repositories.Repository{
		Game: repositories.NewGameRepository(opt),
	}
	service := initService(opt, &repo)

	httpServer := http.NewServer(opt, service)
	httpServer.StartApp()

}

func initService(opt commons.Option, repo *repositories.Repository) *usecases.Services {
	svc := usecases.Services{
		Game: game.NewGameService(opt, repo),
	}
	return &svc
}

func initDatabase(app *appcontext.AppContext, cfg *config.Configuration) (db *gorm.DB, err error) {
	db, err = app.GetDBInstance(appcontext.DBDialectMysql, cfg.Database)
	if err != nil {
		err = fmt.Errorf("error connect to DB MySQL | %v", err)
		return
	}
	logger.Println("Connected to Database successfully")

	return
}
