package main

import (
	"fmt"
	"line-bot/infrastructure/config"
	"line-bot/internal/app/handler/http"
	"line-bot/internal/platform/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"gorm.io/gorm"
)

func initializeApp(cfg *config.Config) (*gin.Engine, error) {
	// ===== Setup Database =====
	// _, err := setupDatabase(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	// =====  Initialize LINE Bot client =====
	bot, err := linebot.New(cfg.Line.ChannelSecret, cfg.Line.ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	// ===== Setup Router =====
	router := gin.New()
	router.Use(gin.Recovery())

	// ===== Setup Routes =====
	http.SetupRoutes(router, bot)

	return router, nil
}

func setupDatabase(cfg *config.Config) (*gorm.DB, error) {
	// ===== Initialize Postgres =====
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port)

	// ===== Connect to Postgres =====
	db, err := database.InitializePostgres(dsn)
	if err != nil {
		return nil, err
	}

	// ===== Run Migrations =====
	if err := database.Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}
