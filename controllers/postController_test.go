
package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/controllers"
	"github.com/stretchr/testify/assert"
)

func TestPostCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/posts", controllers.PostCreate)

	body := map[string]string{
		"Title": "test title",
		"Body":  "test body",
	}

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestPostIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/posts", controllers.PostIndex)

	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostFind(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/posts/:id", controllers.PostFind)

	req, err := http.NewRequest("GET", "/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PUT("/posts/:id", controllers.PostUpdate)

	body := map[string]string{
		"Title": "updated title",
		"Body":  "updated body",
	}

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("PUT", "/posts/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
}

func TestPostDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.DELETE("/posts/:id", controllers.PostDelete)

	req, err := http.NewRequest("DELETE", "/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
}