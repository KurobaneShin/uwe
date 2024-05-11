package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"uwe/types"

	"github.com/google/uuid"
)

type CreateFileUploadRequest struct {
	FileType types.FileType `json:"file_type,omitempty"`
	Mapping  map[string]int `json:"mapping,omitempty"`
}

type CreateFileUploadResponse struct {
	ID uuid.UUID `json:"id"`
}

func HandleCreateFileUpload(w http.ResponseWriter, r *http.Request) error {
	var req CreateFileUploadRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	fileUpload := types.FileUpload{
		ID:      uuid.New(),
		Type:    req.FileType,
		Mapping: req.Mapping,
	}

	resp := CreateFileUploadResponse{
		ID: fileUpload.ID,
	}

	return writeJSON(w, http.StatusCreated, resp)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) error {
	// do some checks
	subs, err := processSubscriptions(r.Body)
	if err != nil {
		return err
	}

	ProcessSubscriptions(subs)
	return nil
}

type Mapping struct {
	Amount     int
	Currency   int
	Period     int
	VAT        int
	StartedAt  int
	CanceledAt int
	ExternalId int
}

func processSubscriptions(r io.Reader) ([]types.Subscription, error) {
	reader := csv.NewReader(r)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		fmt.Println(record)
	}

	return nil, nil
}

func ProcessSubscriptions(subs []types.Subscription) error {
	return nil
}
