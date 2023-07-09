package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"yogi_task/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func SkipOdd(s string) (result string) {
	for i, v := range s {
		if i%2 == 0 {
			result = result + string(v)
		}
	}
	return
}

// EchoBackHandler returns the Body of the message from the Request.
func EchoBackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", "Bad Request")
		} else {
			if len(msg) > 0 {
				c.String(http.StatusOK, "%s", msg)
			} else {
				c.String(http.StatusBadRequest, "%s", "Please add a text on the request body")
			}
		}
	}
}

// ReverseHandler returns the Body of the reverse of the message from the Request.
func ReverseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", "Bad Request")
		} else {
			if len(msg) > 0 {
				c.String(http.StatusOK, "%s", Reverse(string(msg)))
			} else {
				c.String(http.StatusBadRequest, "%s", "Please add a text on the request body")
			}
		}
	}
}

// SkipOddHandler returns the Body of the message from the Request with the odd index characters removed.
func SkipOddHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", "Bad Request")
		} else {
			if len(msg) > 0 {
				c.String(http.StatusOK, "%s", SkipOdd(string(msg)))
			} else {
				c.String(http.StatusBadRequest, "%s", "Please add a text on the request body")
			}
		}
	}
}

// PingHandler replies to a ping request with 'pong'
func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	}
}

// getRoutes sets the routes and their handlers for the HTTP server
func getRoutes(logger *logrus.Logger) *gin.Engine {
	r := gin.New()
	gin.DefaultWriter = logger.Writer()
	gin.DefaultErrorWriter = logger.Writer()

	r.Use(middleware.MWLogRequest(logger))
	r.Use(middleware.MWPanicRecovery(logger))
	r.NoRoute(func(c *gin.Context) { // This is a way to define an "Anonymous" function as a Handler.
		c.String(http.StatusNotFound, "%s", "Not Found")
	})
	r.NoMethod(func(c *gin.Context) { // This is a way to define an "Anonymous" function as a Handler.
		c.String(http.StatusMethodNotAllowed, "%s", "Method not allowed.")
	})

	r.GET("/ping", PingHandler()) // Here we are constructing the Handler by calling the "PingHandler"
	r.GET("/echo", EchoBackHandler())
	r.POST("/reverse", ReverseHandler())
	r.POST("/skip_odd", SkipOddHandler())

	return r
}

// SetLogFormatter sets a log formatter for logrus
func SetLogFormatter(l *logrus.Logger, logFormat string) {
	timestampFormat := "2006-01-02T15:04:05.000Z"
	switch strings.ToLower(logFormat) {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{TimestampFormat: timestampFormat})
	case "text":
		l.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: timestampFormat,
			FullTimestamp:   true,
		})
	default:
		l.Fatalf("Unknown log format '%s'", logFormat)
	}
}

func main() {
	ip := "127.0.0.1"
	port := "1423"

	logger := logrus.Logger{Level: logrus.InfoLevel}
	SetLogFormatter(&logger, "text")

	logger.SetOutput(os.Stdout)

	routes := getRoutes(&logger)
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", ip, port),
		Handler:      routes,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.WithFields(logrus.Fields{
		"addr": server.Addr}).Info("starting server")
	err := server.ListenAndServe()
	if err != nil {
		logger.Print(err)
		return
	}
}
