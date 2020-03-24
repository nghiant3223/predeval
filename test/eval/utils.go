package eval

func newVariableMap(args ...interface{}) map[string]string {
	m := make(map[string]string)
	var key, value string
	for i, a := range args {
		if i%2 == 0 {
			key, _ = a.(string)
		} else {
			value, _ = a.(string)
			m[key] = value
		}
	}
	return m
}
