package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"rest-api/repositories"
	"rest-api/router"
	"rest-api/services"
	"rest-api/sources"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	listenAddr := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))

	db, err := sources.NewPostgresSource(
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		os.Getenv("POSTGRES_PASSWORD"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repositoryManager := repositories.NewRepositoryDBManager(db)
	serviceManager := services.NewServiceManager(repositoryManager)
	appRouter := router.NewRouter(serviceManager)

	if err := fasthttp.ListenAndServe(listenAddr, appRouter.GetHandler()); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}

}

func initConfig() error {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	viper.AddConfigPath("configs")
	viper.SetConfigName(os.Getenv("APP_MODE"))
	return viper.ReadInConfig()
}
