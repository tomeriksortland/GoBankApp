package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJSON(response http.ResponseWriter, status int, value any) error {
	response.WriteHeader(status)
	response.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(response).Encode(value)
}

func makeHTTPHandleFunc(apiFunc apiFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if err := apiFunc(response, request); err != nil {
			WriteJSON(response, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (server *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(server.handleAccount))

	log.Println("JSON API server running on port: ", server.listenAddress)

	http.ListenAndServe(server.listenAddress, router)
}

func (server *APIServer) handleAccount(response http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		return server.handleGetAccount(response, request)
	}
	if request.Method == "POST" {
		return server.handleCreateAccount(response, request)
	}
	if request.Method == "DELETE" {
		return server.handleDeleteAccount(response, request)
	}

	return fmt.Errorf("Method not allowed %s", request.Method)
}

func (s *APIServer) handleGetAccount(response http.ResponseWriter, request *http.Request) error {
	account := NewAccount("Tom", "GG")

	return WriteJSON(response, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(response http.ResponseWriter, request *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(response http.ResponseWriter, request *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(response http.ResponseWriter, request *http.Request) error {
	return nil
}
