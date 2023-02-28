package cmd

import (
	"tiny-url/config"
	"tiny-url/handlers"
	"tiny-url/logics"
	"tiny-url/repository"
	"tiny-url/routes"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the url shortner service",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	dbs := config.LoadConfig(configPath)

	repos := repository.NewRepository(dbs.MysqlConnection, dbs.RedisConnection)
	lgcs := logics.NewLogics(repos)
	handler := handlers.NewHandlers(lgcs)

	e := echo.New()
	e.HideBanner = false

	routes.InitializeGroup(e, handler)
	e.Logger.Fatal(e.Start(":8080"))
}

func init() {
	rootCMD.AddCommand(serveCmd)
}
