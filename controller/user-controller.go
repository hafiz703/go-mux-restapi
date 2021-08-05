package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"golang-mux-api/cache"
	"golang-mux-api/entity"
	"golang-mux-api/errors"
	 
)

type controller struct{}

var (
	// userService service.UserService
	userCache   cache.UserCache
)

type UserController interface {
	GetUserByID(response http.ResponseWriter, request *http.Request)
	GetUsers(response http.ResponseWriter, request *http.Request)
	AddUser(response http.ResponseWriter, request *http.Request)
	DeleteUser(response http.ResponseWriter, request *http.Request)
}

func NewUserController(cache cache.UserCache) UserController {
 
	userCache = cache
	return &controller{}
}

func (*controller) GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	users:= userCache.GetAll()
	if users == nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the users"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func (*controller) GetUserByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	userID := strings.Split(request.URL.Path, "/")[2]
	fmt.Println("userid:",userID)
	var user *entity.User = userCache.Get(userID)
	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "No users found!"})
		return

	}else {
		userCache.Set(userID, user)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(user)
	}
 
}

func (*controller) AddUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := userCache.Validate(&user)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
 
	user.ID = rand.Int63()/10000
	fmt.Println(user.ID)
	userCache.Set(strconv.Itoa(int(user.ID)), &user)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (*controller) DeleteUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	userID := strings.Split(request.URL.Path, "/")[2]
	fmt.Println("userid:",userID)
	var user *entity.User = userCache.Get(userID)
	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "No users found!"})
		return

	}else {
		userCache.Delete(userID)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(userID + " deleted")
	}
 
}
 

