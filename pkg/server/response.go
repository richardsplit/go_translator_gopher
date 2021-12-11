package server

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type Response struct {
	http.ResponseWriter
}

func (r Response) WriteJSON(statusCode int, entity interface{}) error {
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(statusCode)

	encoder := json.NewEncoder(r)
	if err := encoder.Encode(entity); err != nil {
		return errors.Wrap(err, "failed to encode json")
	}
	return nil
}

func (r Response) WriteOctet(statusCode int, entity []byte) error {
	r.Header().Set("Content-Type", "application/octet-stream")
	r.WriteHeader(statusCode)

	_, err := r.Write(entity)
	if err != nil {
		return errors.Wrap(err, "failed to write response")
	}
	return nil
}
