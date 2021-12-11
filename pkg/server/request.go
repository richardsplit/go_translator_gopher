package server

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type Request struct {
	*http.Request
}

func (r Request) ReadJSON(entity interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(entity); err != nil {
		return errors.Wrap(err, "failed to decode json")
	}
	return nil
}
