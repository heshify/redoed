package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
