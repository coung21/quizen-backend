package common

import "github.com/rs/zerolog/log"

type SuccesResponse struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewRestResp(status int, message string, data interface{}) SuccesResponse {
	return SuccesResponse{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}
}

type RestError interface {
	Status() int
	Error() string
	Causes() interface{}
}

type ErrResp struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func (e ErrResp) Status() int {
	return e.ErrStatus
}

func (e ErrResp) Error() string {
	return e.Error()
}

func (e ErrResp) Causes() interface{} {
	return e.ErrCauses
}

func NewRestErr(status int, err string, causes interface{}) RestError {
	log.Error().Stack().Msgf("error: %s, causes: %v", err, causes)
	return ErrResp{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: err,
	}
}
