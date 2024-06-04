package main

import (
	"context"
	"fmt"
	"log"
	"os"
	//"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/javiburn/AsteroidsDB/controllers"
	"github.com/javiburn/AsteroidsDB/configs"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Connect to MongoDB
	configs.Init()

	// Define a route for the GET and POST requests on "/" endpoint
	e.POST("/", controllers.PostAsteroid)
	e.GET("/", controllers.GetAsteroids)

	// Define a route for the GET, PATCH and DELETE requests on any endpoint
	e.GET("/*", controllers.GetAsteroidByID)
	e.PATCH("/*", controllers.UpdateAsteroid)
	e.DELETE("/*", controllers.DeleteAsteroid)
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
