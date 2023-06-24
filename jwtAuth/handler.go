package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("scret_key")

var users = map[string]string{
	"user1":"user1",
	"user2":"user2",
}

type Credential struct{

	username string `json:"username"`
	password string `json:"password"`
}

type Claims struct{
	username string `json:"username"`
	jwt.StandardClaims
}

func logIn(w http.ResponseWriter, r *http.Request){

	var credential Credential
	err:= json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credential.username]

	if !ok || expectedPassword != credential.password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims:= &Claims{
		username: credential.username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name: "token",
			Value: tokenString,
			Expires: expirationTime,
		})
}

func refresh(w http.ResponseWriter, r *http.Request){

}

func home(w http.ResponseWriter, r *http.Request){

}