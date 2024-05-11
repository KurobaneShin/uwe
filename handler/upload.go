package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"uwe/types"
)

const maxFileSize = 10 << 20

func handleUpload(w http.ResponseWriter, r *http.Request) error {
	// do some checks
	subs, err := processSubscriptions(r.Body)
	if err != nil {
		return err
	}

	ProcessSubscriptions(subs)
	return nil
}

func processSubscriptions(r io.Reader) ([]types.Subscription, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	fmt.Println(buf.String())

	return nil, nil
}

func ProcessSubscriptions(subs []types.Subscription) error {
	return nil
}
