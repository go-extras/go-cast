package cast

// Bool will return a bool when `v` is of type bool, or has a method:
//
//	type interface {
//		Bool() (bool, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Bool(v interface{}) (bool, error) {

	switch value := v.(type) {
	case bool:
		return value, nil
	case int8:
		return value != 0, nil
	case int16:
		return value != 0, nil
	case int32:
		return value != 0, nil
	case int64:
		return value != 0, nil
	case uint8:
		return value != 0, nil
	case uint16:
		return value != 0, nil
	case uint32:
		return value != 0, nil
	case uint64:
		return value != 0, nil
	case float32:
		return value != 0, nil
	case float64:
		return value != 0, nil
	case booler:
		return value.Bool()
	default:
		return false, internalCannotCastComplainer{expectedType:"bool", actualType:typeof(value)}
	}
}

// MustBool is like Bool, expect panic()s on an error.
func MustBool(v interface{}) bool {

	x, err := Bool(v)
	if nil != err {
		panic(err)
	}

	return x
}

type booler interface {
	Bool() (bool, error)
}
