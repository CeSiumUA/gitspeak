package models

import (
	"encoding/json"
	"io"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func DeserializeErrorResponseFromBody(reader *io.ReadCloser) (*ErrorMessage, error) {
	errorMessage := ErrorMessage{}
	decoder := json.NewDecoder(*reader)
	err := decoder.Decode(&errorMessage)
	return &errorMessage, err
}

func DeserializeErrorResponseFromString(text string) (*ErrorMessage, error) {
	errorMessage := ErrorMessage{}
	err := json.Unmarshal([]byte(text), &errorMessage)
	return &errorMessage, err
}

func DeserializeErrorResponseFromBytes(text []byte) (*ErrorMessage, error) {
	errorMessage := ErrorMessage{}
	err := json.Unmarshal(text, &errorMessage)
	return &errorMessage, err
}
