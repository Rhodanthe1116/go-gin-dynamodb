package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestSignup(t *testing.T) {
	// Build our expected body
	// user = forms.UserSignup
	newUser := gin.H{
		"name":      "hi",
		"birthday":  "2000-11-16",
		"gender":    "Male",
		"photo_url": "https://example.com",
	}
	// body := strings.NewReader(newUser.Encode())
	jsonByte, _ := json.Marshal(newUser)

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/user/signup", bytes.NewReader(jsonByte))
	req.Header.Add("Content-Type", `application/json`)

	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["name"]

	// Make some assertions on the correctness of the response.
	assert.Equal(t, nil, err)
	assert.Equal(t, true, exists)
	assert.Equal(t, newUser["name"], value)
}

func TestUser(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "hello",
	}

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/", nil)
	req.Header.Add("X-Auth-Secret", `W/"大表格"`)

	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
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
