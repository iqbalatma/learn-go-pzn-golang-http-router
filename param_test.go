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

func TestParam(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		productId := params.ByName("id")
		fmt.Fprint(w, "Product "+productId)
	})

	request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
