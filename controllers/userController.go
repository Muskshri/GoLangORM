package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"gorm_demo/models"
	"gorm_demo/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)
//Development flow : models --> repositoty --> services ---> controller --> main(api creation)
func CreateUser(db *gorm.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		var user models.User
		//unmarshelling
        err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//call service to insert the new records.
		user, err = services.CreateUserService(user, db)
		if err != nil {
			// http error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// success response
		w.WriteHeader(http.StatusCreated)
		// marshaling
		json.NewEncoder(w).Encode(user)

	}

}

func GetUsers(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// if more than one user we need some storage using slice
	 var user []models.User    
     user, err := services.GetUsersService(db)
	 if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	 }
     w.WriteHeader(http.StatusOK)
	//  marshal
	json.NewEncoder(w).Encode(user)
	
	}
}

func GetUserById(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		// id fetch karna h http.request
		var user models.User
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		user, err := services.GetUserByIdService(id, db)
        if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	
	    json.NewEncoder(w).Encode(user)
        
	}
}

func UpdateUser(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		// postman se jo bhi ata h will be recived in request r
        // id ,db ,user -->ye sari cheze chiye service ko call karne ke liye 
		// 
		param:= mux.Vars(r)
		id, _:= strconv.Atoi(param["id"])
		// unmarshall---> json to struct
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
        user, err:= services.UpdateUserService(id, user, db)
		if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	
	    json.NewEncoder(w).Encode(user)
	}
}

func DeleteUser(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		// id , db 
		params:= mux.Vars(r)
		id, _:= strconv.Atoi(params["id"])
		err:= services.DeleteUserService(id, db)
		if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//delete returns a status code of 204 i.e. StatusNoContent
		w.WriteHeader(http.StatusNoContent) 
	}
}
