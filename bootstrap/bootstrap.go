package bootstrap

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/api/repository"
	"boilerplate-api/api/routes"
	"boilerplate-api/api/services"
	"boilerplate-api/api/validators"
	"boilerplate-api/cli"
	"boilerplate-api/docs"
	"boilerplate-api/infrastructure"
	"boilerplate-api/seeds"
	"boilerplate-api/utils"
	"context"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
	validators.Module,
	infrastructure.Module,
	cli.Module,
	seeds.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	env infrastructure.Env,
	logger infrastructure.Logger,
	middlewares middlewares.Middlewares,
	database infrastructure.Database,
	cliApp cli.Application,
	migrations infrastructure.Migrations,
	seeds seeds.Seeds,
) {

	appStop := func(context.Context) error {
		logger.Zap.Info("Stopping Application")
		conn, _ := database.DB.DB()
		_ = conn.Close()
		return nil
	}

	if utils.IsCli() {
		lifecycle.Append(fx.Hook{
			OnStart: func(context.Context) error {
				logger.Zap.Info("Starting boilerplate cli Application")
				logger.Zap.Info("------ 🤖 Boilerplate 🤖 (CLI) ------")
				go cliApp.Start()
				return nil
			},
			OnStop: appStop,
		})

		return
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("------------------------")
			logger.Zap.Info("------ Boilerplate 📺 ------")
			logger.Zap.Info("------------------------")

			go func() {
				if env.Environment != "production" && env.HOST != "" {
					logger.Zap.Info("Setting Swagger Host...")
					docs.SwaggerInfo.Host = env.HOST
				}

				if env.Environment == "production" {
					logger.Zap.Info("Migrating DB schema...")
					migrations.Migrate()
				}
				middlewares.Setup()
				routes.Setup()
				logger.Zap.Info("🌱 seeding data...")
				seeds.Run()
				if env.ServerPort == "" {
					_ = handler.Gin.Run()
				} else {
					_ = handler.Gin.Run(":" + env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: appStop,
	})
}
