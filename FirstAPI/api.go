// Declare the package
package main

// Import the required packages
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Download the git package with the comand "go get ."

/* 1. Declare an struct to save album data
   `json:"id"` is a label, to get exactly this fields names
   at json serialization of the data */

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

/* 2. Create basic data and store it in memory
   This API doesn't have a database conection */

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

/*
3. Create a function that return Http response
  - gin.Context have the request information
  - IndentedJSON serialize the structure en JSON and add it to the http response
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

/*
4. Initialize a router
5. Associate the endpoint name "albums" with the function "getAlbums" name
6. Run the server */

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

/*
7. Create the second endpoint to adds an album using the JSON
received in the request body.
*/
func postAlbums(c *gin.Context) {
	// Declare an empty object of type "album" struct
	var newAlbum album

	// BindJSON deserialize JSON into the struct
	// newAlbum and validates the data.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the memory structure.
	albums = append(albums, newAlbum)
	// Return an http answer
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
 8. Create the endpoint to return the information of an specific album id

It requires to receive the id as parameter.
*/
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
