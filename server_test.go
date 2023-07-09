package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPing(t *testing.T) {
	r := SetUpRouter()
	mockResponse := `pong`
	r.GET("/ping", PingHandler())
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	require.Equal(t, mockResponse, string(responseData))
	require.Equal(t, http.StatusOK, w.Code)
}

func TestEcho(t *testing.T) {
	r := SetUpRouter()
	mockResponse := `ping`
	bodyRequest := `ping`
	b := []byte(bodyRequest)

	r.GET("/echo", EchoBackHandler())
	req, _ := http.NewRequest("GET", "/echo", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEchoWithEmptyBody(t *testing.T) {
	r := SetUpRouter()
	mockResponse := "Please add a text on the request body"

	r.GET("/echo", EchoBackHandler())
	req, _ := http.NewRequest("GET", "/echo", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReverse(t *testing.T) {
	r := SetUpRouter()
	mockResponse := `gnip`
	bodyRequest := `ping`
	b := []byte(bodyRequest)

	r.POST("/reverse", ReverseHandler())
	req, _ := http.NewRequest("POST", "/reverse", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReverseWithEmptyBody(t *testing.T) {
	r := SetUpRouter()
	mockResponse := "Please add a text on the request body"

	r.POST("/reverse", ReverseHandler())
	req, _ := http.NewRequest("POST", "/reverse", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSkipOdd(t *testing.T) {
	r := SetUpRouter()
	mockResponse := `pn`
	bodyRequest := `ping`
	b := []byte(bodyRequest)

	r.POST("/skip_odd", SkipOddHandler())
	req, _ := http.NewRequest("POST", "/skip_odd", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSkipOddWithEmptyBody(t *testing.T) {
	r := SetUpRouter()
	mockResponse := "Please add a text on the request body"

	r.POST("/skip_odd", SkipOddHandler())
	req, _ := http.NewRequest("POST", "/skip_odd", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
