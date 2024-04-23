package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccount(w, r)
	}
	if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("unsupported method")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	account, err := s.storage.GetAccount(idInt)

	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	newAccount := &CreateAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(newAccount)
	if err != nil {
		return fmt.Errorf("failed to decode request: %w", err)
	}

	err = s.storage.CreateAccount(&CreateAccountRequest{Fullname: newAccount.Fullname})

	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return WriteJSON(w, http.StatusCreated, newAccount)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
