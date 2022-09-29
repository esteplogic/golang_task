package main

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"net/http"
)

func TestGetData(t *testing.T) {
	router := SetUpRouter()
 
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getdata", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}