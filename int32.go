package cast

import "math"

// Int32 will return an int32 when `v` is of type int32, int16, int8, uint16, uint8 or has a method:
//
//	type interface {
//		Int32() (int32, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Int16() (int16, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Int8() (int8, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint16() (uint16, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint8() (uint8, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Int32(v interface{}) (int32, error) {

	switch value := v.(type) {
	case int32er:
		return func()(int32, error){
			casted, err := value.Int32()
			if nil != err {
				return 0, err
			}
			return int32(casted), nil
		}()
	case int16er:
		return func()(int32, error){
			casted, err := value.Int16()
			if nil != err {
				return 0, err
			}
			return int32(casted), nil
		}()
	case int8er:
		return func()(int32, error){
			casted, err := value.Int8()
			if nil != err {
				return 0, err
			}
			return int32(casted), nil
		}()
	case uint16er:
		return func()(int32, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return int32(casted), nil
		}()
	case uint8er:
		return func()(int32, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return int32(casted), nil
		}()
	case float32:
		if value < math.MinInt32 || value > math.MaxInt32 {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "overflow"}
		}

		if math.Trunc(float64(value)) != float64(value) {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "decimal part found"}
		}

		return int32(value), nil
	case float64:
		if value < math.MinInt32 || value > math.MaxInt32 {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "overflow"}
		}

		if math.Trunc(value) != value {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "decimal part found"}
		}

		return int32(value), nil
	case int64:
		if value < math.MinInt32 || value > math.MaxInt32 {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "overflow"}
		}
		return int32(value), nil
	case uint32:
		if value > math.MaxInt32 {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "overflow"}
		}
		return int32(value), nil
	case uint64:
		if value > math.MaxInt32 {
			return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value), reason: "overflow"}
		}
		return int32(value), nil
	case int32:
		return int32(value), nil
	case int16:
		return int32(value), nil
	case int8:
		return int32(value), nil
	case uint16:
		return int32(value), nil
	case uint8:
		return int32(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value)}
	}
}

// MustInt32 is like Int32, expect panic()s on an error.
func MustInt32(v interface{}) int32 {

	x, err := Int32(v)
	if nil != err {
		panic(err)
	}

	return x
}

type int32er interface {
	Int32() (int32, error)
}
