package main

import (
	ArtistController "goapi/controller/ArtistController"
	CompanyController "goapi/controller/CompanyController"
	PlaylistController "goapi/controller/PlaylistController"
	SongController "goapi/controller/SongController"
	UserController "goapi/controller/UserController"
	albumController "goapi/controller/albumController"
	"goapi/middleware"
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
		albums.GET("/filter", albumController.FilterGenres)
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

	album := protectedRoutes.Group("/albums")
	{
		album.POST("", albumController.Create)
		album.PUT("/:id", albumController.Update)
		album.DELETE("/:id", albumController.Delete)
	}

	// Companies routes
	company := protectedRoutes.Group("/companies")
	{
		company.POST("", CompanyController.Create)
		company.PUT("/:id", CompanyController.Update)
		company.DELETE("/:id", CompanyController.Delete)
	}

	// Artists routes
	artist := protectedRoutes.Group("/artists")
	{
		artist.POST("", ArtistController.Create)
		artist.PUT("/:id", ArtistController.Update)
		artist.DELETE("/:id", ArtistController.Delete)
	}

	// Playlist routes
	playlists := protectedRoutes.Group("/playlist")
	{
		playlists.POST("", PlaylistController.Create)
		playlists.PUT("/:id", PlaylistController.Update)
		playlists.DELETE("/:id", PlaylistController.Delete)
	}

	// Songs routes
	song := protectedRoutes.Group("/songs")
	{
		song.POST("", SongController.Create)
		song.PUT("/:id", SongController.Update)
		song.DELETE("/:id", SongController.Delete)
	}

	router.Run()
}
