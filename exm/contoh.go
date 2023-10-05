// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
// func awal(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"data": "Hello from Gin-gonic & mongoDB",
// 	})
// }
// func main() {
// 	router := gin.Default()

// 	router.GET("/albums", func(c *gin.Context) {
// 		if albums == nil || len(albums) == 0 {
// c.AbortWithStatus(http.StatusNotFound)
// 		} else {
// 			c.IndentedJSON(http.StatusOK, albums)
// 		}
// 	})

// 	router.GET("/albums/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		for _, a := range albums {
// 			if a.ID == id {
// 				c.IndentedJSON(http.StatusOK, a)
// 				return
// 			}
// 		}
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// 	})

// 	router.POST("/albums", postAlbums)
// 	router.GET("/", awal)
// 	router.Run("localhost:8080")
// }
