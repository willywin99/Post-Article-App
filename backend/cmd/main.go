package main

import (
	"backend/config"
	"backend/internal/router"
	"log"
)

func main() {
	db := config.InitDB()
	config.RunMigrations(db)
	r := router.SetupRouter(db)
	port := config.GetEnv("PORT", "8080")
	log.Printf("✅ Backend running on :%s", port)
	r.Run(":" + port)
}
