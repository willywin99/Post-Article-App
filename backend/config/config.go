package config

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	mymigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv" // 👈 load .env
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetEnv safely reads environment variables with fallback
func GetEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

// InitDB connects to MySQL using GORM
func InitDB() *gorm.DB {
	// 👇 Load environment variables from .env file
	if err := godotenv.Load(); err == nil {
		log.Println("📦 .env file loaded successfully")
	} else {
		log.Println("⚠️  .env file not found, using system environment variables")
	}

	user := GetEnv("DB_USER", "root")
	pass := GetEnv("DB_PASS", "")
	host := GetEnv("DB_HOST", "127.0.0.1")
	port := GetEnv("DB_PORT", "3306")
	name := GetEnv("DB_NAME", "article")

	// ✅ Log the loaded values (excluding password)
	log.Printf("🔧 DB Config — User: %s, Host: %s, Port: %s, Name: %s\n", user, host, port, name)

	var dsn string
	if pass == "" {
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", user, host, port, name)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	}

	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect database: %v", err)
	}

	log.Println("✅ Database connected successfully")
	return db
}

// RunMigrations runs all migrations using golang-migrate
func RunMigrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get db instance: %v", err)
	}
	driver, err := mymigrate.WithInstance(sqlDB, &mymigrate.Config{})
	if err != nil {
		log.Fatalf("mysql driver init failed: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)
	if err != nil {
		log.Printf("no migrations to run or migration init failed: %v", err)
		return
	}
	log.Println("🚀 Running migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("✅ Migrations applied (or none were needed)")
}
