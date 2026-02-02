package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", testFunc)
	http.HandleFunc("/db", dbHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func testFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Hello World!")
	}
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		db, err := connectDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		fmt.Fprintln(w, "✅ Database connection successful")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
