package receiver

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ory-am/google-json-style-response/responder"
	"io"
)

type Receiver interface {
	GetResponse(r io.Reader) (*Response, error)
}

type Response struct {
	ApiVersion string          `json:"apiVersion"`
	Id         uuid.UUID       `json:"id,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Error      responder.Error `json:"error,omitempty"`
}

type receiver struct {
	apiVersion string
}

func New(apiVersion string) Receiver {
	return &receiver{
		apiVersion: apiVersion,
	}
}

func (r *receiver) GetResponse(reader io.Reader) (*Response, error) {
	decoder := json.NewDecoder(reader)
	response := new(Response)
	err := decoder.Decode(response)
	if err != nil {
		return nil, err
	}
	if r.apiVersion != response.ApiVersion {
		return nil, errors.New(fmt.Sprintf("API versions mismatch. Expected %s but got %s", r.apiVersion, response.ApiVersion))
	}

	return response, err
}
