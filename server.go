package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Lover", Artist: "Taylor Swift", Price: 29.99},
	{ID: "2", Title: "Reputation", Artist: "Taylor Swift", Price: 26.99},
	{ID: "3", Title: "Folklore", Artist: "Taylor Swift", Price: 35.99},
	{ID: "4", Title: "Midnights", Artist: "Taylor Swift", Price: 29.99},
	{ID: "5", Title: "1989", Artist: "Taylor Swift", Price: 35.99},
}

// getALbums
func getAlbums(d *gin.Context) {
	d.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	r := gin.Default()

	//when request is made to /ping , func (c *gin context is executed)
	// function takes a gin.Context parameter,
	// which provides information about the incoming request
	// and allows you to send a response.

	//gin.Context object is provided by the Gin web framework
	//and encapsulates information about the incoming HTTP request
	//and provides methods for building the HTTP response.
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.Run("localhost:8080")
}
