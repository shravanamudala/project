package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
    "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shravan/workday/models"
)

type userHandler struct {
	userService UserService
} 

func (uh *userHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var useRequest models.UserRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&useRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = validator.New().Struct(useRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = uh.userService.SaveUser(r.Context(), useRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(useRequest)
}

func (uh *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedFields map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	
	data, err := uh.userService.UpdateUser(r.Context(), int64(id), updatedFields)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (uh *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	result, err := uh.userService.GetByID(r.Context(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	data, err := uh.userService.DeleteUser(r.Context(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	result, err := uh.userService.GetUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func NewUserHandler(router *mux.Router, userService UserService) {
	handler := &userHandler{userService: userService}
	router.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user", handler.SaveUser).Methods("POST")
	router.HandleFunc("/user/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/user/{id}", handler.DeleteUser).Methods("DELETE")
}
