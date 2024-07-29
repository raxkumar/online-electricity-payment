package auth

import (
	"context"
	"encoding/json"
	"errors" // Import the errors package
	"net/http"
	"strings"

	app "com.electricity.online/config"
	"github.com/Nerzal/gocloak/v13"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/micro/micro/v3/service/logger"
)

var (
	Client       *gocloak.GoCloak
	clientId     string
	clientSecret string
	RealmName    string
	keycloakUrl  string
)

func SetClient() {
	clientId = app.GetVal("GO_MICRO_CLIENT_ID")
	clientSecret = app.GetVal("GO_MICRO_CLIENT_SECRET")
	RealmName = app.GetVal("GO_MICRO_REALM_NAME")
	keycloakUrl = app.GetVal("GO_MICRO_KEYCLOAK_URL")
	Client = gocloak.NewClient(keycloakUrl)
}

// GetAccessToken retrieves the access token from the request headers
func GetAccessToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < 1 {
		return "", errors.New("Unauthorized: Missing Authorization header")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", errors.New("Unauthorized: Invalid Authorization header format")
	}

	return authParts[1], nil
}

func Protect(next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := GetAccessToken(r)
		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		rptResult, err := Client.RetrospectToken(context.TODO(), accessToken, clientId, clientSecret, RealmName)
		if err != nil {
			logger.Errorf("Inspection failed: %s", err.Error())
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		istokenvalid := *rptResult.Active
		if !istokenvalid {
			logger.Errorf("Token expired")
			w.WriteHeader(401)
			json.NewEncoder(w).Encode("Token expired")
			return
		}

		// Assuming you have decoded the token
		_, claims, err := Client.DecodeAccessToken(context.TODO(), accessToken, RealmName)
		if err != nil {
			// Handle error
			return
		}

		// Check if the user has the required roles
		if len(roles) > 0 && !hasRequiredRoles(claims, roles) {
			logger.Errorf("Insufficient permissions")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Insufficient permissions")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func hasRequiredRoles(claims *jwt.MapClaims, requiredRoles []string) bool {
	// Check if the token claims are available
	if claims == nil {
		return false
	}

	// Extract the roles from the token claims
	roles, ok := (*claims)["realm_access"].(map[string]interface{})["roles"].([]interface{})
	if !ok {
		return false
	}

	// Check if the user has any of the required roles
	for _, requiredRole := range requiredRoles {
		for _, role := range roles {
			if role == requiredRole {
				logger.Info(role)
				return true
			}
		}
	}

	return false
}
