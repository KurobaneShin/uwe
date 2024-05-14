package handler

import (
	"net/http"
	"uwe/db"
)

type ApiKeyHandler struct {
	db db.DB
}

func NewApiKeyHandler(db db.DB) *UploadHandler {
	return &UploadHandler{
		db: db,
	}
}

func (h *ApiKeyHandler) HandleCreate(w http.ResponseWriter, r *http.Request) error {
	return nil
}
