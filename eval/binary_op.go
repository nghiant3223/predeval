package eval

//// Evaluate for op Or And Eq Neq Le Lte Ge Gte In Div Mul Mod Add Sub
//func EvaluateBinaryOperator(left, right interface{}, op binaryop.BinaryOp) interface{} {
//	assureType(left, right, op)
//
//	if leftInt, ok := left.(int64); ok {
//		rightInt, _ := right.(int64)
//		return evaluateBinaryInt(leftInt, rightInt, op)
//	}
//
//	if leftFloat, ok := left.(float64); ok {
//		rightFloat, _ := right.(float64)
//		return evaluateBinaryFloat(leftFloat, rightFloat, op)
//	}
//
//	if leftString, ok := left.(string); ok {
//		rightString, _ := right.(string)
//		return evaluateBinaryString(leftString, rightString, op)
//	}
//
//	if leftBool, ok := left.(bool); ok {
//		rightBool, _ := right.(bool)
//		return evaluateBinaryBool(leftBool, rightBool, op)
//	}
//
//	panic(errors.InvalidDataType)
//}
//
//func assureType(left, right interface{}, op binaryop.BinaryOp) {
//	if op == binaryop.In {
//		options, _ := right.([]interface{})
//		if len(options) == 0 {
//			return
//		}
//
//		ofSameType := true
//		option0Type := reflect.TypeOf(options[0])
//		for _, opt := range options[1:] {
//			if reflect.TypeOf(opt) != option0Type {
//				ofSameType = false
//			}
//		}
//
//		// If values in right operand of IN op is different from each other
//		if !ofSameType {
//			panic(errors.InvalidDataType)
//		}
//
//		// If left operand of IN type is different from that of right operand
//		if reflect.TypeOf(left) != option0Type {
//			panic(errors.InconsistentOperandDataType)
//		}
//	}
//
//	// For the other operators, left operand and right operand must have the same type
//	if reflect.TypeOf(left) != reflect.TypeOf(right) {
//		panic(errors.InconsistentOperandDataType)
//	}
//}
//
//func evaluateBinaryInt(left, right int64, op binaryop.BinaryOp) interface{} {
//	switch op {
//	case binaryop.Or:
//		return left != 0 || right != 0
//	case binaryop.And:
//		return left != 0 && right != 0
//	case binaryop.Eq:
//		return left == right
//	case binaryop.Neq:
//		return left != right
//	case binaryop.Le:
//		return left < right
//	case binaryop.Ge:
//		return left > right
//	case binaryop.Lte:
//		return left <= right
//	case binaryop.Gte:
//		return left >= right
//	default:
//		return nil
//	}
//}
//
//func evaluateBinaryFloat(left, right float64, op binaryop.BinaryOp) interface{} {
//	switch op {
//	case binaryop.Or:
//		return left != 0 || right != 0
//	case binaryop.And:
//		return left != 0 && right != 0
//	case binaryop.Eq:
//		return left == right
//	default:
//		return nil
//	}
//}
//
//func evaluateBinaryString(left, right string, op binaryop.BinaryOp) interface{} {
//	switch op {
//	case binaryop.Or:
//		return left != "" || right != ""
//	case binaryop.And:
//		return left != "" && right != ""
//	case binaryop.Eq:
//		return left == right
//	default:
//		return nil
//	}
//}
//
//func evaluateBinaryBool(left, right bool, op binaryop.BinaryOp) interface{} {
//	switch op {
//	case binaryop.Or:
//		return left || right
//	case binaryop.And:
//		return left && right
//	case binaryop.Eq:
//		return left == right
//	default:
//		return nil
//	}
//}
