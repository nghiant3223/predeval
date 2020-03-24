package errors

type ListenerError int

const (
	InvalidOperation ListenerError = iota + 1
	InvalidDataType
	SyntaxError
	TypeMismatch
)

var mapErrorToMessage = map[ListenerError]string{
	InvalidOperation: "invalid operation",
	SyntaxError:      "syntax error",
	InvalidDataType:  "invalid data type",
	TypeMismatch:     "type mismatch",
}

func (e ListenerError) Error() string {
	return mapErrorToMessage[e]
}
