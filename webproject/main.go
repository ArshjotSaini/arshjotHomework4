package main

import (
	"fmt"
	"net/http"
	"webproject/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	utils.Dbinit()
	utils.Mapinit()

	// Create a new router
	router := mux.NewRouter()

	// registering the handler and routing
	router.HandleFunc("/", utils.HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/company", utils.CompanyHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/new_company", utils.NewCompanyHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/delete", utils.DeleteHandler).Methods(http.MethodGet)
	router.HandleFunc("/map", utils.MapHandler).Methods(http.MethodGet)

	// Serve static files from the "static" directory
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Start the HTTP server on port 9090 using the router
	fmt.Println("Server is listening on :9090...")
	http.ListenAndServe(":9090", router)
}
