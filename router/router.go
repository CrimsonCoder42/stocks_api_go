// router package is responsible for routing incoming HTTP requests to the correct handler functions.
package router

import (
	"github.com/gorilla/mux"          // Import the gorilla/mux library
	"go-postgres-yt/middleware"       // Import the application's middleware package
)

// Router function initializes a new router and defines routes for the application.
func Router() *mux.Router {
	router := mux.NewRouter()         // Create a new router

	// Define routes and associate them with handler functions
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	
	return router  // Return the configured router
}
