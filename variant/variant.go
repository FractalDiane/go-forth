package variant

import "math"

type Variant interface {
	Add(other Variant) Variant
	Sub(other Variant) Variant
	Mul(other Variant) Variant
	Div(other Variant) Variant
	Mod(other Variant) Variant

	And(other Variant) Variant
	Or(other Variant) Variant
	Xor(other Variant) Variant
	Not() Variant

	Eq(other Variant) Variant
	Ne(other Variant) Variant
	Lt(other Variant) Variant
	Gt(other Variant) Variant
	Le(other Variant) Variant
	Ge(other Variant) Variant
}

type ForthBool bool
type ForthInt int64
type ForthFloat float64
type ForthString string

func (b ForthBool) Add(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		return b || otherCast
	case ForthInt:
		var asInt ForthInt
		if b {
			asInt = 1
		} else {
			asInt = 0
		}

		return asInt + otherCast
	case ForthFloat:
		var asFloat ForthFloat
		if b {
			asFloat = 1.0
		} else {
			asFloat = 0.0
		}

		return asFloat + otherCast
	default:
		panic("test")
	}
}

func (b ForthBool) Sub(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		var asInt ForthInt
		if b {
			asInt = 1
		} else {
			asInt = 0
		}

		return asInt - otherCast
	case ForthFloat:
		var asFloat ForthFloat
		if b {
			asFloat = 1.0
		} else {
			asFloat = 0.0
		}

		return asFloat - otherCast
	default:
		panic("test")
	}
}

func (b ForthBool) Mul(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		return b && otherCast
	case ForthInt:
		var asInt ForthInt
		if b {
			asInt = 1
		} else {
			asInt = 0
		}

		return asInt * otherCast
	case ForthFloat:
		var asFloat ForthFloat
		if b {
			asFloat = 1.0
		} else {
			asFloat = 0.0
		}

		return asFloat * otherCast
	default:
		panic("test")
	}
}

func (b ForthBool) Div(other Variant) Variant {
	panic("test")
}

func (b ForthBool) Mod(other Variant) Variant {
	panic("test")
}

func (b ForthBool) And(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		return b && otherCast
	case ForthInt:
		return b && (otherCast != 0)
	case ForthFloat:
		return b && (otherCast != 0.0)
	default:
		panic("test")
	}
}

func (b ForthBool) Or(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		return b || otherCast
	case ForthInt:
		return b || (otherCast != 0)
	case ForthFloat:
		return b || (otherCast != 0.0)
	default:
		panic("test")
	}
}

func (b ForthBool) Xor(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		return ForthBool(b != otherCast)
	case ForthInt:
		return ForthBool(b != (otherCast != 0))
	case ForthFloat:
		return ForthBool(b != (otherCast != 0.0))
	default:
		panic("test")
	}
}

func (b ForthBool) Not() Variant {
	return !b
}

func (b ForthBool) Eq(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		return ForthBool(b == otherCast)
	case ForthInt:
		return ForthBool(b == (otherCast != 0))
	case ForthFloat:
		return ForthBool(b == (otherCast != 0.0))
	default:
		panic("test")
	}
}

func (b ForthBool) Ne(other Variant) Variant {
	return b.Xor(other)
}

func (b ForthBool) Lt(other Variant) Variant {
	panic("test")
}

func (b ForthBool) Gt(other Variant) Variant {
	panic("test")
}

func (b ForthBool) Le(other Variant) Variant {
	panic("test")
}

func (b ForthBool) Ge(other Variant) Variant {
	panic("test")
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (i ForthInt) Add(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		var result = i
		if otherCast {
			result++
		}

		return result
	case ForthInt:
		return i + otherCast
	case ForthFloat:
		return ForthFloat(i) + otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Sub(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		var result = i
		if otherCast {
			result--
		}

		return result
	case ForthInt:
		return i - otherCast
	case ForthFloat:
		return ForthFloat(i) - otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Mul(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		if otherCast {
			return i
		} else {
			return ForthInt(0)
		}
	case ForthInt:
		return i * otherCast
	case ForthFloat:
		return ForthFloat(i) * otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Div(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i / otherCast
	case ForthFloat:
		return ForthFloat(i) / otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Mod(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i % otherCast
	case ForthFloat:
		return ForthFloat(math.Mod(float64(i), float64(otherCast)))
	default:
		panic("test")
	}
}

func (i ForthInt) And(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i & otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Or(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i | otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Xor(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i ^ otherCast
	default:
		panic("test")
	}
}

func (i ForthInt) Not() Variant {
	return ^i
}

func (i ForthInt) Eq(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i == otherCast)
	case ForthFloat:
		return ForthBool(i == ForthInt(otherCast))
	default:
		panic("test")
	}
}

func (i ForthInt) Ne(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i != otherCast)
	case ForthFloat:
		return ForthBool(i != ForthInt(otherCast))
	default:
		panic("test")
	}
}

func (i ForthInt) Lt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i < otherCast)
	case ForthFloat:
		return ForthBool(i < ForthInt(otherCast))
	default:
		panic("test")
	}
}

func (i ForthInt) Gt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i > otherCast)
	case ForthFloat:
		return ForthBool(i > ForthInt(otherCast))
	default:
		panic("test")
	}
}

func (i ForthInt) Le(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i <= otherCast)
	case ForthFloat:
		return ForthBool(i <= ForthInt(otherCast))
	default:
		panic("test")
	}
}

func (i ForthInt) Ge(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i >= otherCast)
	case ForthFloat:
		return ForthBool(i >= ForthInt(otherCast))
	default:
		panic("test")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (f ForthFloat) Add(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		var result = f
		if otherCast {
			result++
		}

		return result
	case ForthInt:
		return f + ForthFloat(otherCast)
	case ForthFloat:
		return f + otherCast
	default:
		panic("test")
	}
}

func (f ForthFloat) Sub(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		var result = f
		if otherCast {
			result--
		}

		return result
	case ForthInt:
		return f - ForthFloat(otherCast)
	case ForthFloat:
		return f - otherCast
	default:
		panic("test")
	}
}

func (f ForthFloat) Mul(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthBool:
		if otherCast {
			return f
		} else {
			return ForthFloat(0.0)
		}
	case ForthInt:
		return f * ForthFloat(otherCast)
	case ForthFloat:
		return f * otherCast
	default:
		panic("test")
	}
}

func (f ForthFloat) Div(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return f / ForthFloat(otherCast)
	case ForthFloat:
		return f / otherCast
	default:
		panic("test")
	}
}

func (f ForthFloat) Mod(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthFloat(math.Mod(float64(f), float64(otherCast)))
	case ForthFloat:
		return ForthFloat(math.Mod(float64(f), float64(otherCast)))
	default:
		panic("test")
	}
}

func (f ForthFloat) And(other Variant) Variant {
	panic("test")
}

func (f ForthFloat) Or(other Variant) Variant {
	panic("test")
}

func (f ForthFloat) Xor(other Variant) Variant {
	panic("test")
}

func (f ForthFloat) Not() Variant {
	panic("test")
}

func (f ForthFloat) Eq(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f == ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f == otherCast)
	default:
		panic("test")
	}
}

func (f ForthFloat) Ne(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f != ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f != otherCast)
	default:
		panic("test")
	}
}

func (f ForthFloat) Lt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f < ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f < otherCast)
	default:
		panic("test")
	}
}

func (f ForthFloat) Gt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f > ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f > otherCast)
	default:
		panic("test")
	}
}

func (f ForthFloat) Le(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f <= ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f <= otherCast)
	default:
		panic("test")
	}
}

func (f ForthFloat) Ge(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f >= ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f >= otherCast)
	default:
		panic("test")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (s ForthString) Add(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return s + otherCast
	default:
		panic("test")
	}
}

func (s ForthString) Sub(other Variant) Variant {
	panic("test")
}

func (s ForthString) Mul(other Variant) Variant {
	panic("test")
}

func (s ForthString) Div(other Variant) Variant {
	panic("test")
}

func (s ForthString) Mod(other Variant) Variant {
	panic("test")
}

func (s ForthString) And(other Variant) Variant {
	panic("test")
}

func (s ForthString) Or(other Variant) Variant {
	panic("test")
}

func (s ForthString) Xor(other Variant) Variant {
	panic("test")
}

func (s ForthString) Not() Variant {
	panic("test")
}

func (s ForthString) Eq(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s == otherCast)
	default:
		panic("test")
	}
}

func (s ForthString) Ne(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s != otherCast)
	default:
		panic("test")
	}
}

func (s ForthString) Lt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s < otherCast)
	default:
		panic("test")
	}
}

func (s ForthString) Gt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s > otherCast)
	default:
		panic("test")
	}
}

func (s ForthString) Le(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s <= otherCast)
	default:
		panic("test")
	}
}

func (s ForthString) Ge(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s >= otherCast)
	default:
		panic("test")
	}
}
