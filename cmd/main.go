package main

import (
	testTaskObjects "JWTService"
	"JWTService/pkg/database"
	"JWTService/pkg/handler"
	"JWTService/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := database.NewMongoDB("mongodb://localhost:27017")

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := database.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(testTaskObjects.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error: %s", err.Error())
	}
}
