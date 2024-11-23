package main

import (
	"fmt"
	"log"
	"net/http"
	"gorm_demo/controllers"
	"gorm_demo/databases"

	"github.com/gorilla/mux"
)

func main(){
	// initialize database
	db, error := databases.InitDatabase()
	if error != nil{
		fmt.Errorf("Db connection not establised...") //gives error
		log.Fatalf("Db connection not establised...") //stops the program and gives error 
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/api/users", controllers.GetUsers(db)).Methods("GET")
    router.HandleFunc("/api/users/{id}", controllers.GetUserById(db)).Methods("GET")
    router.HandleFunc("/api/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")

	fmt.Println("Server starting on: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	
}