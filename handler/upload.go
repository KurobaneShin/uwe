package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"

	"uwe/db"
	"uwe/types"
)

type CreateFileUploadRequest struct {
	FileType   types.FileType `json:"file_type,omitempty"`
	CustomerId uuid.UUID      `json:"customerId,omitempty"`
	Mapping    map[string]int `json:"mapping,omitempty"`
}

func (r CreateFileUploadRequest) validate() map[string]string {
	errors := map[string]string{}

	if r.CustomerId == types.UUIDZERO {
		errors["customerId"] = "invalid customer id"
	}

	if r.FileType != types.FileTypeSubscription {
		errors["fileType"] = "invalid file type"
	}

	return errors
}

type CreateFileUploadResponse struct {
	ID uuid.UUID `json:"id"`
}

type UploadHandler struct {
	db db.DB
}

func NewUploadHandler(db db.DB) *UploadHandler {
	return &UploadHandler{
		db: db,
	}
}

func (h *UploadHandler) HandleCreateFileUpload(w http.ResponseWriter, r *http.Request) error {
	var req CreateFileUploadRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	if errors := req.validate(); len(errors) != 0 {
		return InvalidRequestData(errors)
	}

	fileUpload := types.FileUpload{
		ID:         uuid.New(),
		Type:       req.FileType,
		CustomerId: req.CustomerId,
		Mapping:    req.Mapping,
	}

	if err := h.db.CreateFileUpload(&fileUpload); err != nil {
		return err
	}

	resp := CreateFileUploadResponse{
		ID: fileUpload.ID,
	}

	return writeJSON(w, http.StatusCreated, resp)
}

func (h *UploadHandler) HandleFileUpload(w http.ResponseWriter, r *http.Request) error {
	fileId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return err
	}

	fileUpload, err := h.db.GetFileUploadById(fileId)
	if err != nil {
		return err
	}

	subs, err := processSubscriptionsUpload(r.Body, fileUpload.Mapping)
	if err != nil {
		return err
	}

	ProcessSubscriptions(subs)
	return nil
}

func processSubscriptionsUpload(r io.Reader, mapping map[string]int) ([]types.Subscription, error) {
	if err := validateSubscriptionMapping(mapping); err != nil {
		return nil, fmt.Errorf("Invalid mapping for subscription file %s", err.Error())
	}

	reader := csv.NewReader(r)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		amount := record[mapping["amount"]]

		_ = amount
		fmt.Println(record)
	}

	return nil, nil
}

func validateSubscriptionMapping(m map[string]int) error {
	keys := []string{
		"amount",
		"currency",
		"period",
		"vat",
		"externalId",
		"startedAt",
		"canceledAt",
	}

	for _, key := range keys {
		if _, ok := m[key]; !ok {
			return fmt.Errorf("%s not in mapping", key)
		}
	}
	return nil
}

func ProcessSubscriptions(subs []types.Subscription) error {
	return nil
}
