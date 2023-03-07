package main

import (
	routes "blogs_api/src"
	"blogs_api/utils/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	addr := "localhost:8080"

	r := gin.New()

	setupRoutes(r)
	setupDatabase()
	startServer(addr, r)
}

func setupRoutes(r *gin.Engine) {
	routes.BlogRoutes(r)
}

func setupDatabase() {
	//sets up database
	database.Db()

}

// startServer to start the server
func startServer(port string, r *gin.Engine) {
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Listen: %s\n", err)
	}

}
