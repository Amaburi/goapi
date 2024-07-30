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
	ID         string   `gorm:"primaryKey" json:"id"`
	Title      string   `gorm:"type:varchar(300)" json:"title"`
	Artist     string   `gorm:"type:varchar(300)" json:"artist"`
	Price      float64  `gorm:"type:float" json:"price"`
	PlayList   PlayList `gorm:"foreignKey:PlayListID" json:"playlist"`
	PlayListID string   `json:"playlist_id"`
}

type Company struct {
	ID            string `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(300)" json:"name"`
	Location      string `gorm:"type:varchar(300)" json:"location"`
	CompanyNumber string `gorm:"type:varchar(20)" json:"company_number"`
}

type Artist struct {
	ID                 string `gorm:"primaryKey" json:"id"`
	Name               string `gorm:"type:varchar(300)" json:"name"`
	Age                int    `gorm:"type:int(11)" json:"age"`
	Address            string `gorm:"type:varchar(300)" json:"address"`
	PhoneNumber        string `gorm:"type:varchar(11)" json:"phone_number"`
	SocialMediaAccount string `gorm:"type:varchar(200)" json:"social_media_account"`
	Achievement        string `gorm:"type:varchar(100)" json:"achievement"`
}

type PlayList struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(300)" json:"name"`
	Artist string `gorm:"type:varchar(300)" json:"artist"`
	Likes  int    `gorm:"type:int" json:"likes"`
	Saved  int    `gorm:"type:int" json:"saved"`
	Songs  []Song `gorm:"foreignKey:PlayListID" json:"songs"`
}

type Song struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(300)" json:"name"`
	Artist     string `gorm:"type:varchar(300)" json:"artist"`
	Duration   string `gorm:"type:varchar(50)" json:"duration"`
	PlayListID string `json:"playlist_id"`
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
	err = database.AutoMigrate(&Album{}, &Company{}, &Artist{}, &PlayList{}, &Song{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	DB = database
}
