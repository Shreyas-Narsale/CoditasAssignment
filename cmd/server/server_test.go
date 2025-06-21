package server

//go test ./cmd/server -v
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func callAPI(t *testing.T, router *gin.Engine, route string, jsonPayload []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", route, bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestTestApi_ValidInput(t *testing.T) {
	router := NewRouter()

	validPayload := []byte(`{
        "name": "Shreyas",
        "pan_card": "ABCWE1234E",
        "mobile": "9876543210",
        "email": "john.doe@example.com"
    }`)

	w := callAPI(t, router, "/v1/api/test", validPayload)

	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"StatusCode":200`)
	assert.Contains(t, body, `"message":"User Validated Successfully"`)
}

func TestTestApi_InvalidJSON(t *testing.T) {
	router := NewRouter()

	invalidJSON := []byte(`{ invalid json }`)

	w := callAPI(t, router, "/v1/api/test", invalidJSON)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"StatusCode":400`)
	assert.Contains(t, body, `"error":"Cannot parse JSON"`)
}

func TestTestApi_ValidationError(t *testing.T) {
	router := NewRouter()

	invalidPayload := []byte(`{
        "name": "",
        "pan_card": "INVALIDPAN",
        "mobile": "123",
        "email": "not-an-email"
    }`)

	w := callAPI(t, router, "/v1/api/test", invalidPayload)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"StatusCode":400`)
	assert.Contains(t, body, `"error":"Validation failed"`)
}

func TestInvalidPANFormat(t *testing.T) {
	router := NewRouter()

	payload := []byte(`{
        "name": "Shreyas",
        "pan_card": "INVALIDPAN123",
        "mobile": "9876543210",
        "email": "john.doe@example.com"
    }`)

	w := callAPI(t, router, "/v1/api/test", payload)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"error":"Validation failed"`)
	assert.Contains(t, body, `"PAN"`)
}

func TestInvalidMobileFormat(t *testing.T) {
	router := NewRouter()

	payload := []byte(`{
        "name": "Shreyas",
        "pan_card": "ABCWE1234E",
        "mobile": "12345abc",
        "email": "john.doe@example.com"
    }`)

	w := callAPI(t, router, "/v1/api/test", payload)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"error":"Validation failed"`)
	assert.Contains(t, body, `"Mobile"`)
}

func TestInvalidEmailFormat(t *testing.T) {
	router := NewRouter()

	payload := []byte(`{
        "name": "Shreyas",
        "pan_card": "ABCWE1234E",
        "mobile": "9876543210",
        "email": "invalid-email-format"
    }`)

	w := callAPI(t, router, "/v1/api/test", payload)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"error":"Validation failed"`)
	assert.Contains(t, body, `"Email"`)
}

func TestMissingFields(t *testing.T) {
	router := NewRouter()

	payload := []byte(`{}`)

	w := callAPI(t, router, "/v1/api/test", payload)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"error":"Validation failed"`)

	assert.Contains(t, body, `"Name"`)
	assert.Contains(t, body, `"PAN"`)
	assert.Contains(t, body, `"Mobile"`)
	assert.Contains(t, body, `"Email"`)
}
