package errors

import (
	"elder-wand/utils/log"
	"fmt"
	"time"
)

type Error struct {
	Uuid   string
	Code   ResponseCode
	Detail string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error uuid [%s] and error code is [%d] and message is [%s]", e.Uuid, e.Code, e.Detail)
}

func New(code ResponseCode, detail string) *Error {
	errMsg := GetMessage(code)
	if code != IgnoreError {
		log.Errorw(fmt.Sprintf("%s: %s", errMsg, detail), "Code", code)
	}
	return &Error{
		Code:   code,
		Detail: detail,
	}
}

func Newf(code ResponseCode, format string, a ...interface{}) *Error {
	detail := fmt.Sprintf(format, a...)
	errMsg := GetMessage(code)
	if code != IgnoreError {
		log.Errorw(fmt.Sprintf("%s: %s", errMsg, detail), "Code", code)
	}
	return &Error{
		Code:   code,
		Detail: detail,
		Uuid:   time.Now().Format("20060102150405000"),
	}
}

func NewDBError(detail string) *Error {
	return New(DBQueryError, detail)
}

func NewDBErrorf(detail string, a ...interface{}) *Error {
	return Newf(DBQueryError, detail, a...)
}

func NewNotFoundError(detail string) *Error {
	return New(NotFoundError, detail)
}

func NewNotFoundErrorf(detail string, a ...interface{}) *Error {
	return Newf(NotFoundError, detail, a...)
}

func NewParamError(detail string) *Error {
	return New(ParamError, detail)
}

func NewParamErrorf(detail string, a ...interface{}) *Error {
	return Newf(ParamError, detail, a...)
}

func NewUnknownError(detail string) *Error {
	return New(UnknownError, detail)
}

func NewUnknownErrorf(detail string, a ...interface{}) *Error {
	return Newf(UnknownError, detail, a...)
}

func NewHttpNotStatusOKError(detail string) *Error {
	return New(HttpNotStatusOKError, detail)
}

func NewHttpNotStatusOKErrorf(detail string, a ...interface{}) *Error {
	return Newf(HttpNotStatusOKError, detail, a...)
}

func NewParamParsingError(detail string) *Error {
	return New(ParamParsingError, detail)
}

func NewParamParsingErrorf(detail string, a ...interface{}) *Error {
	return Newf(ParamParsingError, detail, a...)
}

func NewHttpBusinessCodeError(detail string) *Error {
	return New(HttpBusinessCodeError, detail)
}

func NewHttpBusinessCodeErrorf(detail string, a ...interface{}) *Error {
	return Newf(HttpBusinessCodeError, detail, a...)
}

func NewDBDataError(detail string) *Error {
	return New(DBDataError, detail)
}

func NewDBDataErrorf(detail string, a ...interface{}) *Error {
	return Newf(DBDataError, detail, a...)
}

func NewDBDataConfigError(detail string) *Error {
	return New(DBDataConfigError, detail)
}

func NewDBDataConfigErrorf(detail string, a ...interface{}) *Error {
	return Newf(DBDataConfigError, detail, a...)
}

func NewPaymentError(detail string) *Error {
	return New(PaymentError, detail)
}

func NewPaymentErrorf(detail string, a ...interface{}) *Error {
	return Newf(PaymentError, detail, a...)
}

func NewIgnoreError(detail string) *Error {
	return New(IgnoreError, detail)
}

func NewIgnoreErrorf(detail string, a ...interface{}) *Error {
	return Newf(IgnoreError, detail, a...)
}

func NewBusinessLogicError(detail string) *Error {
	return New(BusinessLogicError, detail)
}

func NewBusinessLogicErrorf(detail string, a ...interface{}) *Error {
	return Newf(BusinessLogicError, detail, a...)
}

func NewUnsupportedError(detail string) *Error {
	return New(UnsupportedError, detail)
}

func NewUnsupportedErrorf(detail string, a ...interface{}) *Error {
	return Newf(UnsupportedError, detail, a...)
}

func NewFileReadError(detail string) *Error {
	return New(FileReadError, detail)
}

func NewFileReadErrorf(detail string, a ...interface{}) *Error {
	return Newf(FileReadError, detail, a...)
}

func NewUndefinedError(detail string) *Error {
	return New(UndefinedError, detail)
}

func NewUndefinedErrorf(detail string, a ...interface{}) *Error {
	return Newf(UndefinedError, detail, a...)
}

func NewValueError(detail string) *Error {
	return New(ValueError, detail)
}

func NewValueErrorf(detail string, a ...interface{}) *Error {
	return Newf(ValueError, detail, a...)
}
