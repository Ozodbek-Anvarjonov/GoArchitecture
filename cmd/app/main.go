package main

import (
	"Architecture/config"
	"Architecture/internal/app"
	"log"
)

// @title My Go API
// @version 1.0
// @description Bu mening Go API misolim
// @host localhost:8080
// @BasePath /
func main() {
	// Config yuklash
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config yuklashda xatolik: %v", err)
	}

	// Ilovani ishga tushirish
	app.Run(cfg)
}
