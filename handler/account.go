package handler

import (
	"net/http"
	"uwe/db"
	"uwe/types"
)

type AccountHandler struct {
	db db.DB
}

func NewAccountHandler(db db.DB) *UploadHandler {
	return &UploadHandler{
		db: db,
	}
}

type CreateAccountRequest struct{}

func (h *AccountHandler) HandleCreate(w http.ResponseWriter, r *http.Request) error {
	account := types.Account{}
	if err := h.db.CreateAccount(&account); err != nil {
		return err
	}

	apiKey := types.NewAPIKey(account.ID, "default")

	if err := h.db.CreateAPIKey(&apiKey); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, account)
}
