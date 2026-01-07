package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gloonch/CarZone/driver"
	carHandler "github.com/gloonch/CarZone/handler/car"
	engineHandler "github.com/gloonch/CarZone/handler/engine"
	loginHandler "github.com/gloonch/CarZone/handler/login"
	"github.com/gloonch/CarZone/middleware"
	carService "github.com/gloonch/CarZone/service/car"
	engineService "github.com/gloonch/CarZone/service/engine"
	carStore "github.com/gloonch/CarZone/store/car"
	engineStore "github.com/gloonch/CarZone/store/engine"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.NewStore(db)
	carService := carService.NewCarService(carStore)

	engineStore := engineStore.NewEngineStore(db)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	schemaFile := "./store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil {
		log.Fatalf("Error while executing the schema file: ", err)
	}

	router.HandleFunc("/login", loginHandler.LoginHandler).Methods("POST")

	// Middleware
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/cars/{id}", carHandler.GetCarByID).Methods("GET")
	protected.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	protected.HandleFunc("/cars", carHandler.GetCarByID).Methods("POST")
	protected.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	protected.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	protected.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("GET")
	protected.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	protected.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	protected.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

}

func executeSchemaFile(db *sql.DB, file string) error {
	sqlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
