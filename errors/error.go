package errors

const BuiltInDivideByZero = "runtime.errorString(\"integer divide by zero\")"

type ListenerError int

const (
	InvalidOperation ListenerError = iota + 1
	InvalidDataType
	SyntaxError
	VariableHasNoValue
	DivideByZero
)

var mapErrorToMessage = map[ListenerError]string{
	InvalidOperation:   "invalid operation",
	SyntaxError:        "syntax error",
	InvalidDataType:    "invalid data type",
	VariableHasNoValue: "variable has no value",
	DivideByZero:       "divide by zero",
}

func (e ListenerError) Error() string {
	return mapErrorToMessage[e]
}
