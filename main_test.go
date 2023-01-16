package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	// setup routes
	router = setupRouter()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestPingRoute(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "pong", rr.Body.String())
}

// func TestUserRoute(t *testing.T) {
// 	rr := httptest.NewRecorder()
// 	req := httptest.NewRequest(http.MethodGet, "/user/abc", nil)
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "no value")
// }
