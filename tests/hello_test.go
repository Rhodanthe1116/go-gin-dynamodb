package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestHello(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "hello",
	}

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["message"]

	// Make some assertions on the correctness of the response.
	assert.Equal(t, nil, err)
	assert.Equal(t, true, exists)
	assert.Equal(t, body["message"], value)
}
