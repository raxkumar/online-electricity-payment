package service

import (
	"encoding/json"
	"net/http"

	"com.electricity.online/models"
	"com.electricity.online/repository"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

var userRepository *repository.UserRepository
var AllowedZoneIDs = []string{"NZ", "SZ", "WZ", "EZ"}

type UserService struct{}

func (us *UserService) AddUser(response http.ResponseWriter, request *http.Request) {
	var user *models.User
	_ = json.NewDecoder(request.Body).Decode(&user)

	// Validate zone_id
	if user.ZoneID != nil && !isValidZoneID(user.ZoneID) {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid zoneID value" }`))
		return
	}
	err := userRepository.CreateUser(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Inserted data")
	json.NewEncoder(response).Encode(user)
}

func (us *UserService) GetAllUsers(response http.ResponseWriter, request *http.Request) {
	users, err := userRepository.GetAllUsers()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func (us *UserService) GetUserByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID := vars["id"]

	user, err := userRepository.GetUserByID(userID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "User not found" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func (us *UserService) UpdateUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID := vars["id"]

	var updatedUser models.User
	_ = json.NewDecoder(request.Body).Decode(&updatedUser)

	// Validate zone_id
	if updatedUser.ZoneID != nil && !isValidZoneID(updatedUser.ZoneID) {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid zoneID value" }`))
		return
	}

	updatedUserDetails, err := userRepository.UpdateUserByID(userID, &updatedUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(updatedUserDetails)
}

func (us *UserService) GetUsersByZone(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	zoneID := vars["zoneId"]

	// Validate zone_id
	if zoneID != "" && !isValidZoneID(&zoneID) {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid zoneID value" }`))
		return
	}

	users, err := userRepository.GetUsersByZone(zoneID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func isValidZoneID(zoneID *string) bool {
	for _, allowedID := range AllowedZoneIDs {
		if *zoneID == allowedID {
			return true
		}
	}
	return false
}
