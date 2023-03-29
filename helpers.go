package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type ApiError struct {
	Error string
}

const jwtSecret = "Super_Secret_WOW"

type apiFunc func(http.ResponseWriter, *http.Request) error

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

			writeJSON(w, http.StatusBadRequest, ApiError{
				Error: err.Error(),
			})
		}
	}
}

func getID(request *http.Request) (int, error) {
	idParam := mux.Vars(request)["id"]
	id, err := strconv.Atoi(idParam) //cast id value from string to int
	if err != nil {
		return id, fmt.Errorf("Invalid id type given %s", idParam)
	}

	return id, nil
}

func withJWTAuth(handlerFunc http.HandlerFunc, storage IStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenAsString := r.Header.Get("x-jwt-token")
		token, err := JwtValidation(tokenAsString)
		if err != nil {
			writeJSON(w, http.StatusForbidden, ApiError{
				Error: "Invalid token",
			})
			return
		}

		if !token.Valid {
			writeJSON(w, http.StatusForbidden, ApiError{
				Error: "Invalid token",
			})
			return
		}

		userId, err := getID(r)
		if err != nil {
			writeJSON(w, http.StatusForbidden, ApiError{
				Error: "access denined",
			})
			return
		}
		account, err := storage.GetAccountByID(userId)
		if err != nil {
			writeJSON(w, http.StatusForbidden, ApiError{
				Error: "access denined",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if account.Number != int64(claims["account_number"].(float64)) {
			writeJSON(w, http.StatusForbidden, ApiError{
				Error: "access denined",
			})
			fmt.Println(account)
			fmt.Println(claims)
			return
		}

		handlerFunc(w, r)
	}

}

func JwtValidation(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

}

func CreateToken(account *Account) (string, error) {

	claims := &jwt.MapClaims{
		"ExpiresAt":      15000,
		"account_number": account.Number,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))

}
