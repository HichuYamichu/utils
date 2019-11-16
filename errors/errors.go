package errors

type Error struct {
	Code int
	Msg  string
}

func New(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

func NewFromErr(code int, msg error) *Error {
	return &Error{Code: code, Msg: msg.Error()}
}
