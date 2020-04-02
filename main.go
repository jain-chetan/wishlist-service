package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jain-chetan/wishlist-service/interfaces"
	"github.com/jain-chetan/wishlist-service/lib/database"
	"github.com/jain-chetan/wishlist-service/model"
	apiServices "github.com/jain-chetan/wishlist-service/recievers"
)

func main() {
	err := initDBClient()
	if err != nil {
		log.Fatal("DB Driver error", err)
	}

	// to initialize all API's handlers
	api := apiServices.Initialization()

	//router initialization
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/wishlist/ping", api.APIHandler.PingHandler).Methods("GET")
	router.HandleFunc("/wishlist/version", api.APIHandler.VersionHandler).Methods("GET")
	router.HandleFunc("/wishlist", api.PostHandler.PostWishlistHandler).Methods("POST")
	router.HandleFunc("/wishlist", api.GetHandler.GetSingleWishlistHandler).Methods("GET")
	router.HandleFunc("/wishlist/{productID}", api.DeleteHandler.DeleteWishlistHandler).Methods("Delete")

	http.ListenAndServe(":8000", router)
}

func initDBClient() error {
	var config model.DBConfig
	//Read DB credentials from environment variables
	config.User = "mongodb"
	config.Port = "27017"
	config.Host = "localhost"
	interfaces.DBClient = new(database.DBRepo)
	err := interfaces.DBClient.DBConnect(config)
	return err
}
