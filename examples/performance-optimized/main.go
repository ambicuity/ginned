package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// For maximum performance, use the optimized constructor
	// This automatically sets:
	// - gin.ReleaseMode (no debug logs, no colorized output)
	// - FastLogger (1000x+ faster than standard logger)
	// - Optimized GOMAXPROCS
	// - Trusted platform configuration for AppEngine
	r := gin.DefaultOptimized()

	// Example of using pre-marshaled JSON responses for maximum speed
	// These avoid JSON marshaling overhead entirely
	r.GET("/ping", func(c *gin.Context) {
		c.FastPong() // Pre-marshaled {"message":"pong"}
	})

	r.GET("/health", func(c *gin.Context) {
		c.FastOk() // Pre-marshaled {"status":"ok"}
	})

	r.GET("/status", func(c *gin.Context) {
		c.FastSuccess() // Pre-marshaled {"status":"success"}
	})

	// For dynamic content, use FastJSON which uses Sonic when available
	r.GET("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		response := gin.H{
			"user_id": userID,
			"status":  "active",
		}
		c.FastJSON(200, response)
	})

	// For frequently requested dynamic content, consider pre-marshaling
	// and caching the JSON response
	popularResponse := []byte(`{"popular":true,"cached":true,"fast":true}`)
	r.GET("/popular", func(c *gin.Context) {
		c.PreMarshaledJSON(200, popularResponse)
	})

	// Example of optimal server configuration
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	log.Printf("Starting optimized Gin server on :8080")
	log.Printf("Build info: %+v", gin.GetBuildInfo())
	log.Fatal(server.ListenAndServe())
}