package response

import (
    "net/http"
    "code.google.com/p/go-uuid/uuid"
)

type Responder struct {
    apiVersion string
    error *ErrorResponse
    data *DataResponse
}

func New(apiVersion string) *Responder {
    return &Responder{
        apiVersion: apiVersion,
        error: &ErrorResponse{
            Error: Error{
                Errors: make([]ErrorItem, 0, 1),
            },
        },
        data: &DataResponse{},
    }
}

func (r *Responder) Success(data interface{}) Response {
     r.data = &DataResponse{
        ApiVersion: r.apiVersion,
        Id: uuid.NewRandom(),
        Data: data,
    }
    return *r.data
}

func (r *Responder) AddError(e ErrorItem) {
    r.error.Error.Errors = append(r.error.Error.Errors, e)
}

func (r *Responder) Error(code int, message string) Response {
    r.error = &ErrorResponse{
        ApiVersion: r.apiVersion,
        Id: uuid.NewRandom(),
        Error: Error{
            Code: code,
            Message: message,
        },
    }
    return r.error
}

func (r *Responder) Write(w http.ResponseWriter, s Response) error {
    m, err := s.Marshal()
    if err != nil {
        return err
    }
    w.Header().Set("Content-Type", "application/json")
    switch s := s.(type) {
        case *ErrorResponse:
            if s.Error.Code == 0 {
                s.Error.Code = http.StatusInternalServerError
            }
            code := s.Error.Code
            w.WriteHeader(code)
    }
    w.Write(m)
    return nil
}
