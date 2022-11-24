package app

import (
	"encoding/json"
	"microservice_go/domain"
	"net/http"
)

func getAllUsersUsecase(w http.ResponseWriter, r *http.Request) {
	users := []domain.User{
		{Name: "Henrique", Email: "henriqueguedesformiga@gmail.com"},
	}

	json.NewEncoder(w).Encode(users)
}
