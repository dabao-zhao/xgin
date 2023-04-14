package xerrors

import "net/http"

type Error struct {
	Code int64
	Msg  string
}

func (e Error) Error() string {
	return e.Msg
}

func New(code int64, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

func TransToCodeAndMsg(err error) (int64, string) {
	xError, ok := err.(Error)
	if ok {
		return xError.Code, xError.Msg
	}
	return http.StatusInternalServerError, err.Error()
}
