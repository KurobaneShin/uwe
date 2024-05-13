package handler

import (
	"encoding/json"
	"net/http"
	"uwe/db"
	"uwe/types"

	chi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CreateCustomerRequest struct {
	Name string
}

func (c CreateCustomerRequest) validate() map[string]string {
	errors := map[string]string{}
	if len(c.Name) < 3 {
		errors["name"] = "name should be at least 3 characteres"
	}
	return errors
}

type CustomerHandler struct {
	db db.DB
}

func NewCustomerHandler(db db.DB) *CustomerHandler {
	return &CustomerHandler{db: db}
}

func (h *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) error {
	var req CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}

	defer r.Body.Close()

	if errors := req.validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}

	customer := types.Customer{
		AccountId: 1,
		Name:      req.Name,
	}

	if err := h.db.CreateCustomer(&customer); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, customer)
}

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
