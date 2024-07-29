package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	auth "com.electricity.online/auth"
	"github.com/Nerzal/gocloak/v13"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gorilla/mux"
)

type IAMService struct{}

// NewUser represents the user data needed for registration
type NewUser struct {
	Username string
	Password string
	Email    string
	// Add other user attributes as needed
}

// UserResponse represents the user information to be returned in the response
type UserResponse struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
}

// handleError handles the error, sets the appropriate status code, logs the error, and writes the response message
func handleError(response http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError

	// Check for specific error messages and adjust the status code accordingly
	if strings.Contains(err.Error(), "403 Forbidden") {
		statusCode = http.StatusForbidden
	} else if strings.Contains(err.Error(), "401 Unauthorized") {
		statusCode = http.StatusUnauthorized
	}

	response.WriteHeader(statusCode)
	logger.Info(err)
	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
}

// CreateUser creates a new user in Keycloak
func (iamservice *IAMService) CreateUser(response http.ResponseWriter, request *http.Request) {
	accessToken, _ := auth.GetAccessToken(request)
	// Decode the request body to get the NewUser data
	var newUser NewUser
	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid request body")
		return
	}

	user := gocloak.User{
		Username: &newUser.Username,
		Email:    &newUser.Email,
		Enabled:  gocloak.BoolP(true),
		Attributes: &map[string][]string{
			// You can add custom attributes if needed
			"customAttribute": {"value"},
		},
	}

	createdUserID, err := auth.Client.CreateUser(context.TODO(), accessToken, auth.RealmName, user)
	if err != nil {
		handleError(response, err)
		return
	}

	createdUser, err := auth.Client.GetUserByID(context.TODO(), accessToken, auth.RealmName, createdUserID)
	if err != nil {
		handleError(response, err)
		return
	}

	// Handle the success case
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(createdUser)
}

// ListUsers returns a list of all users (excluding admin) from Keycloak
func (iamservice *IAMService) ListUsers(response http.ResponseWriter, request *http.Request) {
	accessToken, _ := auth.GetAccessToken(request)

	// Get the list of all users from Keycloak
	users, err := auth.Client.GetUsers(context.TODO(), accessToken, auth.RealmName, gocloak.GetUsersParams{})
	if err != nil {
		handleError(response, err)
		return
	}

	// Filter out the admin user
	var filteredUsers []UserResponse
	for _, user := range users {
		if user.Username != nil && *user.Username != "admin" {
			filteredUsers = append(filteredUsers, UserResponse{
				Username:      *user.Username,
				Email:         *user.Email,
				EmailVerified: *user.EmailVerified,
			})
		}
	}

	// Handle the success case
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(filteredUsers)
}

// GetUserByID retrieves a user by ID from Keycloak
func (iamservice *IAMService) GetUserByID(response http.ResponseWriter, request *http.Request) {
	accessToken, _ := auth.GetAccessToken(request)
	vars := mux.Vars(request)
	userID, ok := vars["id"]
	if !ok {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("User ID is required in the URL")
		return
	}
	user, err := auth.Client.GetUserByID(context.TODO(), accessToken, auth.RealmName, userID)
	if err != nil {
		handleError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

// UpdateUserDetails updates a user's details in Keycloak
func (iamservice *IAMService) UpdateUserDetails(response http.ResponseWriter, request *http.Request) {
	accessToken, _ := auth.GetAccessToken(request)
	// Extract user ID from the request URL parameters
	vars := mux.Vars(request)
	userID, ok := vars["id"]
	if !ok {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("User ID is required in the URL")
		return
	}

	// Decode the request body to get the updated user details
	var updatedUser gocloak.User
	err := json.NewDecoder(request.Body).Decode(&updatedUser)
	updatedUser.ID = &userID
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid request body")
		return
	}

	// Update the user in Keycloak
	err = auth.Client.UpdateUser(context.TODO(), accessToken, auth.RealmName, updatedUser)
	if err != nil {
		handleError(response, err)
		return
	}

	// Handle the success case
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("User details updated successfully")
}

// DisableUser disables a user in Keycloak
func (iamservice *IAMService) DisableUser(response http.ResponseWriter, request *http.Request) {
	accessToken, _ := auth.GetAccessToken(request)

	// Extract user ID from the request URL parameters
	vars := mux.Vars(request)
	userID, ok := vars["id"]
	if !ok {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("User ID is required in the URL")
		return
	}

	// Disable the user in Keycloak
	err := auth.Client.UpdateUser(context.TODO(), accessToken, auth.RealmName, gocloak.User{
		ID:      &userID,
		Enabled: gocloak.BoolP(false),
	})
	if err != nil {
		handleError(response, err)
		return
	}

	// Handle the success case
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("User disabled successfully")
}
