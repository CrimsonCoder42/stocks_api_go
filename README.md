<!-- README.md -->

# Go-Postgres Stock Management API

This project provides a simple API for managing stock data, built with Go and PostgreSQL. Below is a detailed breakdown of the core components of this API.

## Table of Contents
- [Project Structure](#project-structure)
- [Models](#models)
- [Router](#router)
- [Middleware Handlers](#middleware-handlers)
- [Getting Started](#getting-started)

## Project Structure
├── main.go
├── models
│ └── models.go
├── router
│ └── router.go
└── middleware
└── handlers.go


## Models
The `models` package (found in `models/models.go`) contains the data structures that are used within the application.

### Stock Struct
- `StockID`: Unique identifier for the stock (integer).
- `Name`: Name of the stock (string).
- `Price`: Price of the stock in cents (integer), represented as an int64 to avoid floating-point imprecision.
- `Company`: Company to which the stock belongs (string).

## Router
The `router` package (found in `router/router.go`) is responsible for routing incoming HTTP requests to the appropriate handler functions. It utilizes the `gorilla/mux` library to efficiently manage routes.

- Routes are defined and associated with handler functions from the `middleware` package.
- The router supports various HTTP methods for each route, including `GET`, `POST`, `PUT`, `DELETE`, and `OPTIONS`.

## Middleware Handlers
The `middleware` package contains handler functions for processing incoming requests and interacting with the PostgreSQL database.

- `createConnection()`: Establishes and returns a connection to the database.
- `CreateStock()`: Handles the creation of a new stock in the database.
- `GetStock()`: Retrieves a single stock from the database.
- `GetAllStocks()`: Retrieves all stocks from the database.
- `UpdateStock()`: Updates a stock in the database.
- `DeleteStock()`: Deletes a stock from the database.
- Additional utility functions to insert, update, delete, and retrieve stocks.

## Getting Started
1. Ensure that Go and PostgreSQL are installed on your system.
2. Clone this repository.
3. Navigate to the project directory and run `go run main.go` to start the server.
4. The API will be accessible at `http://localhost:8080`.

## License
This project is open source and available under the [MIT License](LICENSE).

