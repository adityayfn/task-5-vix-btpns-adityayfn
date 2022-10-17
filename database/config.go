package database

import (
	"fmt"
	"os"

	"github.com/adityayfn/task-5-vix-btpns-adityayfn/app"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Setup Database connection
func SetupDbConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&app.Photo{}, &app.User{})
	// db.AutoMigrate(&app.Photo{}, &app.User{})
	println("Database connected!")
	return db
}
// CloseDbConnection 
func CloseDbConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}