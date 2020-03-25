package eval

import (
	"github.com/ntncsebku/reloader/ast/expression"
	"github.com/ntncsebku/reloader/errors"
)

func (v *visitor) isZeroValue(value interface{}) (bool, error) {
	switch value.(type) {
	case int64:
		intValue := value.(int64)
		return intValue == expression.ZeroIntLiteral.Value, nil
	case float64:
		floatValue := value.(float64)
		return floatValue == expression.ZeroFloatLiteral.Value, nil
	case bool:
		boolValue := value.(bool)
		return boolValue == expression.ZeroBoolLiteral.Value, nil
	case string:
		stringValue := value.(string)
		return stringValue == expression.ZeroStringLiteral.Value, nil
	default:
		return false, errors.InvalidDataType
	}
}
