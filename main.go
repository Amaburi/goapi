package main

import (
	ArtistController "goapi/controller/ArtistController"
	CompanyController "goapi/controller/CompanyController"
	PlaylistController "goapi/controller/PlaylistController"
	SongController "goapi/controller/SongController"
	UserController "goapi/controller/UserController"
	albumController "goapi/controller/albumController"
	"goapi/middleware" // Import the middleware package
	models "goapi/model/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	router.POST("/api/login", UserController.Login)
	router.POST("/api/register", UserController.CreateUser)
	// Public routes
	albums := router.Group("/api/albums")
	{
		albums.GET("/", albumController.Index)
		albums.GET("/:id", albumController.Show)
	}

	companies := router.Group("/api/companies")
	{
		companies.GET("/", CompanyController.Index)
		companies.GET("/:id", CompanyController.Show)
	}

	artists := router.Group("/api/artists")
	{
		artists.GET("/", ArtistController.Index)
		artists.GET("/:id", ArtistController.Show)
	}

	playlist := router.Group("/api/playlist")
	{
		playlist.GET("/", PlaylistController.Index)
		playlist.GET("/:id", PlaylistController.Show)
	}

	songs := router.Group("/api/songs")
	{
		songs.GET("/", SongController.Index)
		songs.GET("/:id", SongController.Show)
	}

	// Protected routes
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		albums.POST("/albums", albumController.Create)
		albums.PUT("/albums/:id", albumController.Update)
		albums.DELETE("/albums/:id", albumController.Delete)

		companies.POST("/companies", CompanyController.Create)
		companies.PUT("/companies/:id", CompanyController.Update)
		companies.DELETE("/companies/:id", CompanyController.Delete)

		artists.POST("/artists", ArtistController.Create)
		artists.PUT("/artists/:id", ArtistController.Update)
		artists.DELETE("/artists/:id", ArtistController.Delete)

		playlist.POST("/playlist", PlaylistController.Create)
		playlist.PUT("/playlist/:id", PlaylistController.Update)
		playlist.DELETE("/playlist/:id", PlaylistController.Delete)

		songs.POST("/songs", SongController.Create)
		songs.PUT("/songs/:id", SongController.Update)
		songs.DELETE("/songs/:id", SongController.Delete)
	}

	router.Run()
}
