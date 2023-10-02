// main package is the entry point of the application.
package main

import (
	"fmt"
	"go-postgres-yt/router"  // Import the application's router package
	"log"
	"net/http"
)

// main function initializes the application and starts the server.
func main() {
	r := router.Router()  // Initialize a new router
	fmt.Println("Starting server on the port 8080...")  // Log that the server is starting
	
	// Start the server on port 8080 and log any errors that occur
	log.Fatal(http.ListenAndServe(":8080", r))
}
