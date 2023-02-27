package tools

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secret = "my_super_secret"

func GenerateToken(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	expTime := time.Now().Add(10 * time.Minute)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expTime.Unix()
	claims["user_id"] = userId

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func ExtractData(tokenString string) (bool, int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot parse token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, -1, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := claims["user_id"].(float64)
		return true, int(userId), nil
	}
	return false, -1, nil

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
