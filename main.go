package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
/*
curl http://localhost:8080/albums \
    --header \
    "Content-Type: application/json" \
    --request "GET"
*/

func postAlbums(c *gin.Context){
	var newAlbum album

	if err:= c.BindJSON(&newAlbum); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, albums)
		return 
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}
/*
curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: application/json" \
    --request "POST" --data \
    '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
*/

func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return 
		}
	}        
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context){
	id := c.Param("id")

	for i,a := range albums {
		if a.ID == id {
			albums = append(albums[:i],  albums[i+1:]...)
            c.Status(http.StatusNoContent)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No album found with given ID"})
}
/*
curl -X DELETE http://localhost:8080/albums/1
*/

func updateAlbumByID(c *gin.Context){
	id := c.Param("id")
	var updatedAlbum album

	if err:= c.BindJSON(&updatedAlbum); err!=nil{
		c.IndentedJSON(http.StatusBadRequest, albums)
		return 
	}

	for i,a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
/*
curl -X PUT http://localhost:8080/albums/1 \
	--include --header \
	"Content-Type: application/json" \
	--request "PUT" --data \
	'{"id": "1","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
*/

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)

	router.Run("localhost:8080")
}