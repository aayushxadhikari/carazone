package main

import (
	"carazone/driver"
	"carazone/service/car_service"
	"carazone/service/engine_service"
	"carazone/store/car_store"
	"carazone/store/engine_store"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"carazone/handler/car_handler"
	"carazone/handler/engine_handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()

	carStore := car_store.New(db)
	carService := car_service.NewCarService(carStore)

	engineStore := engine_store.New(db)
	engineService := engine_service.NewEngineService(engineStore)

	carHandler := car_handler.NewCarHandler(carService)
	engineHandler := engine_handler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	schemaFile := "store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err!= nil{
		log.Fatal("Error while executing the schema file:", err)
	}

	router.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func executeSchemaFile(db *sql.DB, fileName string) error {
	sqlFile, err := os.ReadFile(fileName)
	if err != nil{
		return err
	}

	if _, err := db.Exec(string(sqlFile)); err != nil{
		return err
	}
	return nil
}	
