package errors

type ListenerError int

const (
	InvalidOperation ListenerError = iota + 1
	SyntaxError
)

var mapErrorToMessage = map[ListenerError]string{
	InvalidOperation: "invalid operation",
	SyntaxError:      "syntax error",
}

func (e ListenerError) Error() string {
	return mapErrorToMessage[e]
}
