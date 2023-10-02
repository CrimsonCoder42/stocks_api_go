package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres-yt/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// response struct defines the structure for API responses.
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// createConnection initializes and returns a connection to the database.
func createConnection() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the PostgreSQL database using environment variable
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

// CreateStock handles the creation of a new stock in the database.
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	// Decode the incoming JSON payload into 'stock' struct
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// Insert stock into database and get the inserted ID
	insertID := insertStock(stock)

	// Build and send the API response
	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// GetStock handles the retrieval of a single stock from the database.
func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Convert ID from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// Retrieve stock from the database
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}

	// Send the retrieved stock as a response
	json.NewEncoder(w).Encode(stock)
}

// GetAllStocks handles the retrieval of all stocks from the database.
func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatalf("Unable to get all stocks. %v", err)
	}
	// Send the retrieved stocks as a response
	json.NewEncoder(w).Encode(stocks)
}

// UpdateStock handles the update of a stock in the database.
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Convert ID from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	var stock models.Stock

	// Decode the incoming JSON payload into 'stock' struct
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// Update stock in the database and get the number of rows affected
	updatedRows := updateStock(int64(id), stock)

	// Build and send the API response
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

// DeleteStock handles the deletion of a stock from the database.
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Convert ID from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// Delete stock from the database and get the number of rows affected
	deletedRows := deleteStock(int64(id))

	// Build and send the API response
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

// insertStock inserts a new stock into the database and returns its ID.
func insertStock(stock models.Stock) int64 {
	// Create a connection to the database
	db := createConnection()
	defer db.Close()

	// SQL query to insert a new stock
	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

// getStock retrieves a stock from the database using its ID.
func getStock(id int64) (models.Stock, error) {
	// Create a connection to the database
	db := createConnection()
	defer db.Close()

	var stock models.Stock

	// SQL query to retrieve a stock by ID
	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlStatement, id)

	// Scan the retrieved row into the 'stock' struct
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return stock, err
}

// getAllStocks retrieves all stocks from the database.
func getAllStocks() ([]models.Stock, error) {
	// Create a connection to the database
	db := createConnection()
	defer db.Close()

	var stocks []models.Stock

	// SQL query to retrieve all stocks
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	// Iterate through the result set and scan each row into a 'stock' struct
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

// updateStock updates a stock in the database and returns the number of rows affected.
func updateStock(id int64, stock models.Stock) int64 {
	// Create a connection to the database
	db := createConnection()
	defer db.Close()

	// SQL query to update a stock by ID
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Get the number of rows affected by the update
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

// deleteStock deletes a stock from the database and returns the number of rows affected.
func deleteStock(id int64) int64 {
	// Create a connection to the database
	db := createConnection()
	defer db.Close()

	// SQL query to delete a stock by ID
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Get the number of rows affected by the deletion
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

