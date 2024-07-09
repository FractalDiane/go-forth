package variant

import (
	"log"
	"math"
)

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

	AsBool() bool
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
		log.Fatalf("Error: Invalid '+' operands (%v and %v)", b, other)
		return nil
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
		log.Fatalf("Error: Invalid '-' operands (%v and %v)", b, other)
		return nil
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
		log.Fatalf("Error: Invalid '*' operands (%v and %v)", b, other)
		return nil
	}
}

func (b ForthBool) Div(other Variant) Variant {
	log.Fatalf("Error: Invalid '/' operands (%v and %v)", b, other)
	return nil
}

func (b ForthBool) Mod(other Variant) Variant {
	log.Fatalf("Error: Invalid '%%' operands (%v and %v)", b, other)
	return nil
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
		log.Fatalf("Error: Invalid 'and' operands (%v and %v)", b, other)
		return nil
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
		log.Fatalf("Error: Invalid 'or' operands (%v and %v)", b, other)
		return nil
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
		log.Fatalf("Error: Invalid 'xor' operands (%v and %v)", b, other)
		return nil
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
		log.Fatalf("Error: Invalid '==' operands (%v and %v)", b, other)
		return nil
	}
}

func (b ForthBool) Ne(other Variant) Variant {
	return b.Xor(other)
}

func (b ForthBool) Lt(other Variant) Variant {
	log.Fatalf("Error: Invalid '<' operands (%v and %v)", b, other)
	return nil
}

func (b ForthBool) Gt(other Variant) Variant {
	log.Fatalf("Error: Invalid '>' operands (%v and %v)", b, other)
	return nil
}

func (b ForthBool) Le(other Variant) Variant {
	log.Fatalf("Error: Invalid '<=' operands (%v and %v)", b, other)
	return nil
}

func (b ForthBool) Ge(other Variant) Variant {
	log.Fatalf("Error: Invalid '>=' operands (%v and %v)", b, other)
	return nil
}

func (b ForthBool) AsBool() bool {
	return bool(b)
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
		log.Fatalf("Error: Invalid '+' operands (%v and %v)", i, other)
		return nil
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
		log.Fatalf("Error: Invalid '-' operands (%v and %v)", i, other)
		return nil
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
		log.Fatalf("Error: Invalid '*' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Div(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i / otherCast
	case ForthFloat:
		return ForthFloat(i) / otherCast
	default:
		log.Fatalf("Error: Invalid '/' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Mod(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i % otherCast
	case ForthFloat:
		return ForthFloat(math.Mod(float64(i), float64(otherCast)))
	default:
		log.Fatalf("Error: Invalid '%%' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) And(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i & otherCast
	default:
		log.Fatalf("Error: Invalid 'and' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Or(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i | otherCast
	default:
		log.Fatalf("Error: Invalid 'or' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Xor(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return i ^ otherCast
	default:
		log.Fatalf("Error: Invalid 'xor' operands (%v and %v)", i, other)
		return nil
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
		log.Fatalf("Error: Invalid '==' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Ne(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i != otherCast)
	case ForthFloat:
		return ForthBool(i != ForthInt(otherCast))
	default:
		log.Fatalf("Error: Invalid '!=' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Lt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i < otherCast)
	case ForthFloat:
		return ForthBool(i < ForthInt(otherCast))
	default:
		log.Fatalf("Error: Invalid '<' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Gt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i > otherCast)
	case ForthFloat:
		return ForthBool(i > ForthInt(otherCast))
	default:
		log.Fatalf("Error: Invalid '>' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Le(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i <= otherCast)
	case ForthFloat:
		return ForthBool(i <= ForthInt(otherCast))
	default:
		log.Fatalf("Error: Invalid '<=' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) Ge(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(i >= otherCast)
	case ForthFloat:
		return ForthBool(i >= ForthInt(otherCast))
	default:
		log.Fatalf("Error: Invalid '>=' operands (%v and %v)", i, other)
		return nil
	}
}

func (i ForthInt) AsBool() bool {
	return i != 0
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
		log.Fatalf("Error: Invalid '+' operands (%v and %v)", f, other)
		return nil
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
		log.Fatalf("Error: Invalid '-' operands (%v and %v)", f, other)
		return nil
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
		log.Fatalf("Error: Invalid '*' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Div(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return f / ForthFloat(otherCast)
	case ForthFloat:
		return f / otherCast
	default:
		log.Fatalf("Error: Invalid '/' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Mod(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthFloat(math.Mod(float64(f), float64(otherCast)))
	case ForthFloat:
		return ForthFloat(math.Mod(float64(f), float64(otherCast)))
	default:
		log.Fatalf("Error: Invalid '%%' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) And(other Variant) Variant {
	log.Fatalf("Error: Invalid 'and' operands (%v and %v)", f, other)
	return nil
}

func (f ForthFloat) Or(other Variant) Variant {
	log.Fatalf("Error: Invalid 'or' operands (%v and %v)", f, other)
	return nil
}

func (f ForthFloat) Xor(other Variant) Variant {
	log.Fatalf("Error: Invalid 'xor' operands (%v and %v)", f, other)
	return nil
}

func (f ForthFloat) Not() Variant {
	log.Fatalf("Error: Invalid 'not' operand (%v)", f)
	return nil
}

func (f ForthFloat) Eq(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f == ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f == otherCast)
	default:
		log.Fatalf("Error: Invalid '==' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Ne(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f != ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f != otherCast)
	default:
		log.Fatalf("Error: Invalid '!=' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Lt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f < ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f < otherCast)
	default:
		log.Fatalf("Error: Invalid '<' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Gt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f > ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f > otherCast)
	default:
		log.Fatalf("Error: Invalid '>' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Le(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f <= ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f <= otherCast)
	default:
		log.Fatalf("Error: Invalid '<=' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) Ge(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthInt:
		return ForthBool(f >= ForthFloat(otherCast))
	case ForthFloat:
		return ForthBool(f >= otherCast)
	default:
		log.Fatalf("Error: Invalid '>=' operands (%v and %v)", f, other)
		return nil
	}
}

func (f ForthFloat) AsBool() bool {
	return f != 0.0
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (s ForthString) Add(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return s + otherCast
	default:
		log.Fatalf("Error: Invalid '+' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) Sub(other Variant) Variant {
	log.Fatalf("Error: Invalid '-' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) Mul(other Variant) Variant {
	log.Fatalf("Error: Invalid '*' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) Div(other Variant) Variant {
	log.Fatalf("Error: Invalid '/' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) Mod(other Variant) Variant {
	log.Fatalf("Error: Invalid '%%' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) And(other Variant) Variant {
	log.Fatalf("Error: Invalid 'and' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) Or(other Variant) Variant {
	log.Fatalf("Error: Invalid 'or' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) Xor(other Variant) Variant {
	log.Fatalf("Error: Invalid 'xor' operands (%v and %v)", s, other)
	return nil
}

func (s ForthString) Not() Variant {
	log.Fatalf("Error: Invalid 'not' operand (%v)", s)
	return nil
}

func (s ForthString) Eq(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s == otherCast)
	default:
		log.Fatalf("Error: Invalid '==' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) Ne(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s != otherCast)
	default:
		log.Fatalf("Error: Invalid '!=' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) Lt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s < otherCast)
	default:
		log.Fatalf("Error: Invalid '<' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) Gt(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s > otherCast)
	default:
		log.Fatalf("Error: Invalid '>' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) Le(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s <= otherCast)
	default:
		log.Fatalf("Error: Invalid '<=' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) Ge(other Variant) Variant {
	switch otherCast := other.(type) {
	case ForthString:
		return ForthBool(s >= otherCast)
	default:
		log.Fatalf("Error: Invalid '>=' operands (%v and %v)", s, other)
		return nil
	}
}

func (s ForthString) AsBool() bool {
	return len(s) != 0
}
