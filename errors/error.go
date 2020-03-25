package errors

const BuiltInDivideByZero = "runtime.errorString(\"integer divide by zero\")"

type ReloadError int

const (
	InvalidOperation ReloadError = iota + 1
	InvalidDataType
	SyntaxError
	VariableHasNoValue
	DivideByZero
)

var mapErrorToMessage = map[ReloadError]string{
	InvalidOperation:   "invalid operation",
	SyntaxError:        "syntax error",
	InvalidDataType:    "invalid data type",
	VariableHasNoValue: "variable has no value",
	DivideByZero:       "divide by zero",
}

func (e ReloadError) Error() string {
	return mapErrorToMessage[e]
}
