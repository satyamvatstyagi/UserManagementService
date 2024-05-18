package jwt

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/env"
)

// Function to generate jwt token
func GenerateToken(userID string, createdAt time.Time) (string, error) {
	// Get the jwt secret from the environment variable
	jwtSecret := env.EnvConfig.JWTSecretKey
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET not set")
	}

	// Get the jwt expiry from the environment variable
	jwtExpiry := env.EnvConfig.JWTExpirationTime
	if jwtExpiry == "" {
		return "", fmt.Errorf("JWT_EXPIRY not set")
	}

	// Convert the jwt expiry to int64
	jwtExpiryInt, err := strconv.ParseInt(jwtExpiry, 10, 64)
	if err != nil {
		return "", err
	}

	// Set the claims for the token
	claims := make(jwt.MapClaims)
	claims["https://mymtn.com/loginCount"] = 6
	claims["https://mymtn.com/secondaryProfile"] = []interface{}{
		map[string]interface{}{
			"user_id":    userID,
			"provider":   "sms",
			"connection": "sms",
			"isSocial":   false,
		},
		map[string]interface{}{
			"profileData": map[string]interface{}{
				"email":           "Firstname.Lastname@mtn.com",
				"email_verified":  "true",
				"nonce_supported": true,
				"first_name":      "Firstname",
				"last_name":       "Lastname",
				"name":            "Firstname Lastname",
			},
			"user_id":       userID,
			"refresh_token": "re1f76a0bdaa3437aac406791e4f56aa3.0.mrtsu.b-I-9K7ssi1sZ5_RQoUfcw",
			"provider":      "apple",
			"connection":    "apple",
			"isSocial":      true,
		},
	}
	claims["https://mymtn.com/phone_number"] = "+26876000000"
	claims["https://mymtn.com/country"] = "SWZ"
	claims["https://skillsacademy.mtn.com/connection"] = "sms"
	claims["https://skillsacademy.mtn.com/email"] = "876060571@skillsacademy.mtn.com"
	claims["nickname"] = "+26876000000"
	claims["name"] = "+26876000000"
	claims["picture"] = "testPicture.png"
	claims["updated_at"] = "2023-12-21T13:38:13.917Z"
	claims["iss"] = "localhost"                        // Issuer of the token
	claims["aud"] = "testValue"                        // Audience of the token
	claims["iat"] = createdAt                          // Issued at time of the token
	claims["sub"] = userID                             // Subject of the token
	claims["auth_time"] = time.Now().Unix()            // Time of authentication
	claims["sid"] = "fVWiwb4QMmU565YGZw3HQQsT-Dfa_2S7" // Session ID of the token

	// Set the expiration time for the token
	claims["exp"] = time.Now().Add(time.Duration(jwtExpiryInt) * time.Minute).Unix() // Expiration time of the token

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function to validate jwt token
func ValidateToken(tokenString string) error {
	// Get the jwt secret from the environment variable
	jwtSecret := env.EnvConfig.JWTSecretKey
	if jwtSecret == "" {
		return fmt.Errorf("jwt secret not set")
	}

	// Remove the Bearer prefix from the token if it exists
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check the signing method
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return fmt.Errorf("token error: %v", err)
	}

	// Validate the token
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// Function to get the claims from the token
func GetClaims(tokenString string) (jwt.MapClaims, error) {
	// Get the jwt secret from the environment variable
	jwtSecret := env.EnvConfig.JWTSecretKey
	if jwtSecret == "" {
		return nil, fmt.Errorf("jwt secret not set")
	}

	// Remove the Bearer prefix from the token if it exists
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check the signing method
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token error: %v", err)
	}

	// Validate the token
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
