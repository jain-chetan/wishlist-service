package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jain-chetan/wishlist-service/interfaces"
	"github.com/jain-chetan/wishlist-service/lib/database"
	"github.com/jain-chetan/wishlist-service/model"
)

func main() {
	err := initDBClient()
	if err != nil {
		log.Fatal("DB Driver error", err)
	}
	//router initialization
	mux := http.NewServeMux()
	//Simple ping API
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingResponse := model.Response{
			Code:    200,
			Message: "OK",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(pingResponse)
	})
	http.ListenAndServe(":8000", mux)
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

//Main defines the point where execution starts
// func main() {
// 	err := initDBClient()
// 	if err != nil {
// 		log.Fatal("DB Driver error", err)
// 	}
// 	e := echo.New()

// 	/*if err != nil {
// 		log.Fatal("Unable to proceed due to " + err.Error())
// 	}*/

// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{"*"},
// 	}))
// 	//api := apiServices.Initialization()
// 	g := e.Group("/cart") // added to setup initial group of api's
// 	g.GET("/ping", api.Get.PingHandler)
// 	g.GET("/", api.Get.GetCartHandler)
// 	e.Logger.Fatal(e.Start(":8081"))
// }

// func initDBClient() error {

// 	var config model.DBConfig

// 	config.User = "mongodb"
// 	config.Port = "27017"
// 	config.Host = "localhost"
// 	interfaces.DBClient = new(database.DBRepo)
// 	err := interfaces.DBClient.DBConnect(config)
// 	return err
// }
