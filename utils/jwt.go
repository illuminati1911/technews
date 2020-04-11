package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/illuminati1911/technews/models"
)

// GenerateJWTforUser generates token for regular user
func GenerateJWTforUser(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.UserID,
		"username": user.Username,
		"isAdmin":  user.IsAdmin,
		"expires":  strconv.FormatInt(time.Now().Add(time.Minute*time.Duration(1)).Unix(), 10),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// IsJWTValid verifies JWT and checks that it has not expired.
func IsJWTValid(token string) bool {
	verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := verifiedToken.Claims.(jwt.MapClaims); ok && verifiedToken.Valid {
		parsedExpires, err := strconv.ParseInt(fmt.Sprintf("%v", claims["expires"]), 10, 64)
		if err == nil {
			return parsedExpires > time.Now().Unix()
		}
	}
	return false
}
