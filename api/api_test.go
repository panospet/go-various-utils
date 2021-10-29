package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_HealthCheck(t *testing.T) {
	api, err := New()
	assert.Nil(t, err)
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	api.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"hello!\"\n", w.Body.String())
}
