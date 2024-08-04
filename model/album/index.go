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
	ID          uint     `gorm:"primaryKey" json:"id"`
	Title       string   `gorm:"size:300" json:"title"`
	Artist      string   `gorm:"size:300" json:"artist"`
	Price       float64  `json:"price"`
	PlayList    PlayList `gorm:"foreignKey:PlayListID"`
	PlayListID  uint     `json:"playlist_id"`
	Description string   `gorm:"size:300" json:"description"`
	Awards      string   `gorm:"size:300" json:"awards"`
	Genre       string   `gorm:"size:300" json:"genre"`
	Relasedate  string   `gorm:"size:300" json:"releasedate"`
	Rating      float64  `json:"rating"`
	Link        string   `gorm:"size:300" json:"link"`
	CoverArt    string   `gorm:"size:300" json:"cover_art"`
}

type Company struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"size:300" json:"name"`
	Location      string `gorm:"size:300" json:"location"`
	CompanyNumber string `gorm:"size:20" json:"company_number"`
}

type Artist struct {
	ID                 uint    `gorm:"primaryKey" json:"id"`
	Name               string  `gorm:"size:300" json:"name"`
	Age                int     `json:"age"`
	Address            string  `gorm:"size:300" json:"address"`
	PhoneNumber        string  `gorm:"size:11" json:"phone"`
	SocialMediaAccount string  `gorm:"size:200" json:"social_media_account"`
	Achievement        string  `gorm:"size:100" json:"achievement"`
	Biography          string  `gorm:"size:1000" json:"biography"`
	Nationality        string  `gorm:"size:100" json:"nationality"`
	Website            string  `gorm:"size:200" json:"website"`
	Email              string  `gorm:"size:100" json:"email"`
	Labels             Company `gorm:"foreignKey:CompanyID" json:"labels"`
	CompanyID          uint    `json:"company_id"`
	Albums             []Album `gorm:"many2many:artist_albums" json:"albums"`
	ProfilePicture     string  `gorm:"size:255" json:"profile_picture"`
}

type PlayList struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:300" json:"name"`
	Artist string `gorm:"size:300" json:"artist"`
	Likes  uint   `json:"likes"`
	Saved  uint   `json:"saved"`
	Songs  []Song `gorm:"many2many:playlist_songs" json:"songs"`
}

type PlaylistSongs struct {
	PlayListID uint `gorm:"primaryKey"`
	SongID     uint `gorm:"primaryKey"`
}

type Song struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:300" json:"name"`
	Artist     string `gorm:"size:300" json:"artist"`
	Duration   string `gorm:"size:300" json:"duration"`
	CoverArt   string `gorm:"size:300" json:"cover_art"`
	Relasedate string `gorm:"size:300" json:"releasedate"`
	Genre      string `gorm:"size:300" json:"genre"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// AutoMigrate the models
	database.AutoMigrate(&Album{}, &Company{}, &Artist{}, &PlayList{}, &Song{}, &PlaylistSongs{}, User{})
	DB = database
}
