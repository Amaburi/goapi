package main

import (
	ArtistController "goapi/controller/ArtistController"
	CompanyController "goapi/controller/CompanyController"
	albumController "goapi/controller/albumController"
	models "goapi/model/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	albums := router.Group("/api/albums")
	{
		albums.GET("/", albumController.Index)
		albums.GET("/:id", albumController.Show)
		albums.POST("/", albumController.Create)
		albums.PUT("/:id", albumController.Update)
		albums.DELETE("/:id", albumController.Delete)
	}
	companies := router.Group("/api/companies")
	{
		companies.GET("/", CompanyController.Index)
		companies.GET("/:id", CompanyController.Show)
		companies.POST("/", CompanyController.Create)
		companies.PUT("/:id", CompanyController.Update)
		companies.DELETE("/:id", CompanyController.Delete)
	}
	artists := router.Group("/api/artists")
	{
		artists.GET("/", ArtistController.Index)
		artists.GET("/:id", ArtistController.Show)
		artists.POST("/", ArtistController.Create)
		artists.PUT("/:id", ArtistController.Update)
		artists.DELETE("/:id", ArtistController.Delete)
	}
	router.Run()
}
