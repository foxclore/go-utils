package everything

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer | constraints.Signed
}

func Abs[N Number](num N) N {
	if num < 0 {
		return num
	} else {
		return num
	}
}

// Abs2 was before generics
func Abs2(num interface{}) float64 {
	var res float64
	var multiplier float64 = 1
	switch v := num.(type) {
	// For int we just convert it to float64
	case int:
		res = float64(v)
	case int8:
		res = float64(v)
	case int16:
		res = float64(v)
	case int32:
		res = float64(v)
	case int64:
		res = float64(v)
	// If for some reason we got uint - we can safely return it, those are never negative
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		res = float64(v)
	case float64:
		res = v // Hey, they are the same type!
	default:
		return 0 // Sadly no idea what they gave us
	}

	if res < 0 {
		multiplier = -1
	}

	return multiplier * res
}
