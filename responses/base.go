package responses

import (
	"elder-wand/errors"
)

type JSONResponse struct {
	Code    errors.ResponseCode
	Message string
	Data    interface{}
}
