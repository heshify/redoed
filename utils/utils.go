package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/heshify/redoed/internal/models"
)

func ValidateDocument(doc models.Document) error {
	if doc.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateUser(user models.User) error {
	if user.Name == "" {
		return fmt.Errorf("Name is required")
	}

	if user.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if !validEmail(user.Email) {
		return fmt.Errorf("Invalid email")
	}

	return nil
}

func ValidateLoginPayload(credentials models.AuthUser) error {
	if credentials.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if credentials.Password == "" {
		return fmt.Errorf("Password is required")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ParseJSON(r *http.Request, data any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(data); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	if decoder.More() {
		return fmt.Errorf("unexpected extra data in request body")
	}

	return nil
}

func WriteError(w http.ResponseWriter, status int, err error) {
	if err == nil {
		err = errors.New("unknown error")
	}
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
