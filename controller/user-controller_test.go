package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang-mux-api/cache"
	"golang-mux-api/entity"
)
 

const (
	ID    int64  = 123
	NAME string = "USER 1"
	AGE  int64 = 55
)

var (
 
	userCh       cache.UserCache           = cache.NewRedisCache("localhost:6379", 0, 10)
	userController UserController            = NewUserController(userCh)
)

func TestAddUser(t *testing.T) {
	// Create new HTTP request
	var jsonStr = []byte(`{"name":"` + NAME + `","age":` + strconv.Itoa(int(AGE)) + `}`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(userController.AddUser)

	// Record the HTTP Response
	response := httptest.NewRecorder()
	
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var user entity.User
	json.NewDecoder(io.Reader(response.Body)).Decode(&user)

	// Assert HTTP response
	assert.NotNil(t, user.ID)
	assert.Equal(t, NAME, user.Name)
	assert.Equal(t, AGE, user.Age)

	// Cleanup database
	tearDown(user.ID)
}

func setup() {
	var user entity.User = entity.User{
		ID:    ID,
		Name: NAME,
		Age:  AGE,
	}
	 
	userCh.Set(strconv.Itoa(int(user.ID)),&user)
}

func tearDown(userID int64) {
 
	 
	userCh.Delete(strconv.Itoa(int(userID)))
}

func TestGetUsers(t *testing.T) {

	// Insert new user
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/users", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(userController.GetUsers)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var users []entity.User
	json.NewDecoder(io.Reader(response.Body)).Decode(&users)

	// Assert HTTP response
	assert.Equal(t, ID, users[0].ID)
	assert.Equal(t, NAME, users[0].Name)
	assert.Equal(t, AGE, users[0].Age)

	// Cleanup database
	tearDown(ID)
}

func TestGetUserByID(t *testing.T) {

	// Insert new user
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/users/123", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(userController.GetUserByID)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var user entity.User
	json.NewDecoder(io.Reader(response.Body)).Decode(&user)

	// Assert HTTP response
	assert.Equal(t, ID, user.ID)
	assert.Equal(t, NAME, user.Name)
	assert.Equal(t, AGE, user.Age)

	// Cleanup database
	tearDown(ID)
}


func TestDeleteUser(t *testing.T) {

	// Insert new user
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("DELETE", "/users/123", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(userController.DeleteUser)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var resp string
	json.NewDecoder(io.Reader(response.Body)).Decode(&resp)

	// Assert User Deleted
	assert.Equal(t, resp,  strconv.Itoa(int(ID))+" deleted")
	 

	// Cleanup database
	tearDown(ID)
}
