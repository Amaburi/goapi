package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Album struct {
	ID     string  `gorm:"primaryKey" json:"id"`
	Title  string  `gorm:"varchar(300)" json:"title"`
	Artist string  `gorm:"varchar(300)" json:"artist"`
	Price  float64 `gorm:"float" json:"price"`
}
type Company struct {
	ID            string `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"varchar(300)" json:"name"`
	Location      string `gorm:"varchar(300)" json:"location"`
	CompanyNumber string `gorm:"varchar(20)" json:"companynumber"`
}

type Artist struct {
	ID                 string `gorm:"primaryKey" json:"id"`
	Name               string `gorm:"varchar(300)" json:"name"`
	Age                int    `gorm:"int(11)" json:"age"`
	Address            string `gorm:"varchar(300)" json:"address"`
	PhoneNumber        string `gorm:"varchar(11)" json:"phone"`
	SocialMediaAccount string `gorm:"varchar(200)" json:"socialmediaAccount"`
	Achievement        string `gorm:"varchar(100)" json:"achievement"`
}

func ConnectDB() {
	eerr := godotenv.Load()
	if eerr != nil {
		log.Fatalf("Failed to load .env")
	}
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// AutoMigrate will create the "albums" table if it doesn't exist based on the struct definition
	err = database.AutoMigrate(&Album{}, &Company{}, &Artist{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	DB = database
}
