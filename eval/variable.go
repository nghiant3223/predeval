package eval

import (
	"strconv"
	"strings"

	"github.com/ntncsebku/reloader/errors"
)

func (v *visitor) evaluateVariable(name string) interface{} {
	value, ok := v.mapVariableNameToValue[strings.ToLower(name)]
	if !ok {
		panic(errors.VariableHasNoValue)
	}

	return parsingVariableValue(value)
}

func parsingVariableValue(value string) interface{} {
	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return intValue
	}

	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return floatValue
	}

	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue
	}

	return value
}
