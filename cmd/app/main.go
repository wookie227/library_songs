package main

import (
	"os"
	"song_library/internal/database"
	"song_library/internal/server"
	"song_library/internal/services"
	handlers "song_library/internal/transport/rest"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := database.NewDb(database.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USER_NAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialized db: %s", err.Error())
	}

	repos := database.NewRepository(db)
	services := services.NewService(repos)
	hanlers := handlers.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), hanlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
