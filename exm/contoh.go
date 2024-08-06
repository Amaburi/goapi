//func FilterGenres(c *gin.Context) {
//	var album []models.Album

//	if err := c.ShouldBindJSON(&album); err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body: " + err.Error()})
//		return
//	}

//	genre := c.Query("genre")
//	if genre == "" {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "There's no such a genre that you're looking for "})
//		return
//	}

//	result := models.DB.Where("genre = ?", genre).Find(&album)
//	if result.Error != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update album: " + result.Error.Error()})
//		return
//	}
//	if result.RowsAffected == 0 {
//		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
//		return
//	}
// }
