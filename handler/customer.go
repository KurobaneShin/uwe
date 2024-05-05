package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"uwe/types"
)

func HandleGetCustomer(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return NewApiError(http.StatusBadRequest, err)
	}

	customer := types.Customer{
		Id: id,
	}

	return writeJSON(w, http.StatusOK, customer)
}
