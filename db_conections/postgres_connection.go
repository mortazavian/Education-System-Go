package db_conections

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"net/url"
	"time"
)

func NewPostgres01() (db *sql.DB) {

	// once.Do(func() {
	// Database credentials
	tehranTimezone, _ := time.LoadLocation("Asia/Tehran")
	dbURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("postgres", "admin"),
		Host:     "localhost",
		Path:     "education_system_go",
		RawQuery: "sslmode=disable&timezone=" + tehranTimezone.String(),
	}

	// Convert URL to connection string
	connStr := dbURL.String()

	fmt.Println(connStr)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Execute a test query
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to execute test query:", err)
		return
	}

	fmt.Println("Successfully connected to the database!")
	// })

	return
}
