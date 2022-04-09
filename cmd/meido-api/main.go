package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/meido-api/db"
	"github.com/intrntsrfr/meido-api/handler"
	"github.com/intrntsrfr/meido-api/service"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	DBLogin string `json:"db_login"`
}

func main() {
	// create logger
	logger, _ := zap.NewProduction()
	logger.Info("hello world")

	// read config file
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("config file not found")
	}
	var config *Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic("mangled config file, fix it")
	}

	// create database connection
	dbConn, err := sqlx.Connect("postgres", config.DBLogin)
	if err != nil {
		log.Fatal(err)
	}
	psql := db.NewPsqlDB(dbConn, logger.Named("database"))

	// create services
	jwtUtil := service.NewJWTUtil()
	authService := service.NewAuthService(jwtUtil)

	r := gin.Default()

	cfg := &handler.Config{
		R:           r,
		DB:          psql,
		JWTService:  jwtUtil,
		AuthService: authService,
		Log:         logger.Named("handler"),
	}

	handler.NewHandler(cfg)

	err = r.Run(":4444")
	if err != nil {
		log.Println(err)
		return
	}
	//r.RunTLS(":4444", "./certs/server.crt", "./certs/server.key")
}
