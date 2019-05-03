package controllers

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretKey = "TODO: this is the secret key"
)

type Token struct {
	uid   string
	email string
	// TODO: add more fields
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (t *Token) generateJSONToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":     t.uid,
		"email":   t.email,
		"isAdmin": t.isAdmin,
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
		return ""
	}
	return tokenString
}

func parseJSONToken(tokenString string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error during token decryption")
		}
		return []byte(secretKey), nil
	})
	var data = Token{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data.uid = claims["uid"].(string)
		data.email = claims["email"].(string)
		data.isAdmin = claims["isAdmin"].(bool)
	}
	return &data, err
}
