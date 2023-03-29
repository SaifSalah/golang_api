package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenAddr string
	IStorage   IStorage
}

func newAPIServer(listenAddr string, storage IStorage) *APIServer {

	return &APIServer{
		ListenAddr: listenAddr,
		IStorage:   storage,
	}
}

func (s *APIServer) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountByID), s.IStorage))
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))
	log.Println("JSON API Server running on port: ", s.ListenAddr)
	http.ListenAndServe(s.ListenAddr, router)

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allwed")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	accounts, err := s.IStorage.GetAccounts()

	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {

		id, err := getID(r)

		if err != nil {
			return err
		}

		account, err := s.IStorage.GetAccountByID(id)
		if err != nil {
			return err
		}

		return writeJSON(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)

	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	AccountRequest := new(CreateAccountRequestDto)

	if err := json.NewDecoder(r.Body).Decode(AccountRequest); err != nil {
		return err
	}
	account := NewAccount(AccountRequest.FirstName, AccountRequest.LastName)

	if err := s.IStorage.CreateAccount(account); err != nil {
		return err
	}

	token, err := CreateToken(account)
	if err != nil {
		return err
	}
	fmt.Println(token)

	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	id, err := getID(r)

	if err != nil {
		return err
	}
	if err := s.IStorage.DeleteAccount(id); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]int{"account deleted successfully": id})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	transferDTO := new(TransferAccountRequestDto)

	if err := json.NewDecoder(r.Body).Decode(transferDTO); err != nil {
		return err
	}
	fmt.Println(transferDTO)
	return writeJSON(w, http.StatusOK, transferDTO)
}
