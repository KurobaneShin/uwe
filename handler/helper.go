package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"status_code,omitempty"`
	Msg        any `json:"msg,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewApiError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Msg:        errors,
	}
}

func UnAuthorized() APIError {
	return NewApiError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}

func InvalidJSON() APIError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("invalid JSON Request data"))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"msg":        "Internal server error",
				}
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error())
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
