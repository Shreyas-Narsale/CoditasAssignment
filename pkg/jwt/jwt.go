package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT token for a given user ID, and expiration time
func GenerateJWT(userID string, jwtSecret []byte, expirationTime time.Duration) (string, error) {

	if len(jwtSecret) == 0 {
		return "", errors.New("jwtSecret cannot be empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userID,
		"exp":    time.Now().Add(expirationTime).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWT verifies a given JWT token string and returns the user ID if valid
func VerifyJWT(tokenString string, jwtSecret []byte) (string, error) {

	if len(jwtSecret) == 0 {
		return "", errors.New("jwtSecret cannot be empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("userid not found in token")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
