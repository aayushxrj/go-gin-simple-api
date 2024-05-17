package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)


func router() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	return router
}

func Test_getAlbums(t *testing.T){

	req, err := http.NewRequest("GET", "/albums", nil)
	if err != nil {
		t.Errorf("Error creating request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v\n", http.StatusOK, w.Code)
	}
}

func Test_postAlbums(t *testing.T){

	newAlbum := album{ID: "4", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99}
	jsonValue,_ := json.Marshal(newAlbum)

	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("Error creating request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v\n", http.StatusOK, w.Code)
	}
}
func Test_getAlbumByID(t *testing.T){

	req, err := http.NewRequest("GET", "/albums/1", nil)
	if err != nil {
		t.Errorf("Error creating request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v\n", http.StatusOK, w.Code)
	}
}
func Test_deleteAlbumByID(t *testing.T){

	req, err := http.NewRequest("DELETE", "/albums/1", nil)
	if err != nil {
		t.Errorf("Error creating request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status %v; got %v\n", http.StatusNoContent, w.Code)
	}
}
func Test_updateAlbumByID(t *testing.T){
	
	newAlbum := album{ID: "1", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99}
	jsonValue,_ := json.Marshal(newAlbum)

	req, err := http.NewRequest("PUT", "/albums/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("Error creating request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %v; got %v\n", http.StatusOK, w.Code)
	}
}