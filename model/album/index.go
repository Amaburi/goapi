package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "root:@tcp(localhost:3306)/album" // Replace with your actual credentials
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// AutoMigrate will create the "albums" table if it doesn't exist based on the struct definition
	database.AutoMigrate(&Album{})
	database.AutoMigrate(&Company{})
	database.AutoMigrate(&Artist{})
	DB = database
}
