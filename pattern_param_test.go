package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPatternParam(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:productId/items/:itemId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		productId := params.ByName("productId")
		itemId := params.ByName("itemId")
		fmt.Fprint(w, "Product "+productId+" Item "+itemId)
	})

	request := httptest.NewRequest(http.MethodGet, "/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestPatternCatchAllParam(t *testing.T) {
	router := httprouter.New()

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		fmt.Fprint(w, "Image "+image)
	})

	request := httptest.NewRequest(http.MethodGet, "/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image /small/profile.png", string(body))
}
