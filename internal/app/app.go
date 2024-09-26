package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/auth_test/internal/handler"
	"github.com/auth_test/internal/repository"
	"github.com/auth_test/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RunApp() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("cannot read config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		//Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
		// Host:     os.Getenv("DB_HOST"),
		// Port:     os.Getenv("DB_PORT"),
		// Username: os.Getenv("DB_USERNAME"),
		// //Password: viper.GetString("db.password"),
		// DBName:   os.Getenv("DB_DBNAME"),
		// SSLMode:  os.Getenv("DB_SSLMODE"),
		// Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init DB: ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(Server)
	go func() {
		if err = server.Run(handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running server: %s", err.Error())
		}
	}()
	//logrus.Println("Starting server...")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error while shutting down server: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Fatalf("error while closing db connection: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
