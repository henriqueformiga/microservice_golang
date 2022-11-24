package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// Users Controller
	router.HandleFunc("/users", getAllUsersUsecase).Methods("GET")

	http.ListenAndServe("localhost:8000", nil)
}
