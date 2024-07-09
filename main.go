package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"unicode"

	"goforth/stack"
	"goforth/variant"
)

type variantStack = stack.Stack[variant.Variant]

///////////////////////////////////////////////////////////////////////////////////////////////////

func add(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Add(rhs)
}

func sub(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Sub(rhs)
}

func mul(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Mul(rhs)
}

func div(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Div(rhs)
}

func mod(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Mod(rhs)
}

func and(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.And(rhs)
}

func or(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Or(rhs)
}

func xor(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Xor(rhs)
}

func not(op variant.Variant) variant.Variant {
	return op.Not()
}

func eq(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Eq(rhs)
}

func ne(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Ne(rhs)
}

func lt(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Lt(rhs)
}

func gt(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Gt(rhs)
}

func le(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Le(rhs)
}

func ge(lhs variant.Variant, rhs variant.Variant) variant.Variant {
	return lhs.Ge(rhs)
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func printTop(stack *variantStack) {
	var top = stack.Top()
	fmt.Printf("%v", *top)
	stack.Pop()
}

func printTopLn(stack *variantStack) {
	var top = stack.Top()
	fmt.Printf("%v\n", *top)
	stack.Pop()
}

func emitTop(stack *variantStack) {
	var top = stack.Top()
	switch topCast := (*top).(type) {
	case variant.ForthInt:
		fmt.Printf("%c", rune(topCast))
		stack.Pop()
	default:
		panic("test")
	}
}

func drop(stack *variantStack) {
	stack.Pop()
}

func dup(stack *variantStack) {
	var top = *stack.Top()
	stack.Push(top)
}

func swap(stack *variantStack) {
	stack.SwapTopElements()
}

func over(stack *variantStack) {
	var second = *stack.Second()
	stack.Push(second)
}

func rotate(stack *variantStack) {
	stack.RotateTopElements()
}

func random(stack *variantStack) {
	var value = rand.Int64()
	stack.Push(variant.ForthInt(value))
}

func randomf(stack *variantStack) {
	var value = rand.Float64()
	stack.Push(variant.ForthFloat(value))
}

///////////////////////////////////////////////////////////////////////////////////////////////////

var binaryOperators = map[string]func(variant.Variant, variant.Variant) variant.Variant{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
	"%": mod,

	"==": eq,
	"!=": ne,
	"<":  lt,
	">":  gt,
	"<=": le,
	">=": ge,

	"and": and,
	"or":  or,
	"xor": xor,
}

var unaryOperators = map[string]func(variant.Variant) variant.Variant{
	"not": not,
}

var builtinFunctions = map[string]func(*variantStack){
	".":     printTop,
	",":     printTopLn,
	"emit":  emitTop,
	"drop":  drop,
	"swap":  swap,
	"dup":   dup,
	"over":  over,
	"rot":   rotate,
	"rand":  random,
	"randf": randomf,
}

var forthStack = stack.Stack[variant.Variant]{}
var definedWords = make(map[string][]string, 5)

func executeWord(word string) {
	var wordLower = strings.ToLower(word)
	if integer, err := strconv.Atoi(word); err == nil {
		forthStack.Push(variant.ForthInt(integer))
	} else if float, err := strconv.ParseFloat(word, 64); err == nil {
		forthStack.Push(variant.ForthFloat(float))
	} else if strings.HasPrefix(word, `"`) && strings.HasSuffix(word, `"`) {
		var str = strings.TrimPrefix(word, `"`)
		str = strings.TrimSuffix(str, `"`)
		forthStack.Push(variant.ForthString(str))
	} else if binOpFunction, found := binaryOperators[wordLower]; found {
		var rhs = *forthStack.Top()
		forthStack.Pop()
		var lhs = *forthStack.Top()
		forthStack.Pop()
		forthStack.Push(binOpFunction(lhs, rhs))
	} else if unOpFunction, found := unaryOperators[wordLower]; found {
		var operand = *forthStack.Top()
		forthStack.Pop()
		forthStack.Push(unOpFunction(operand))
	} else if builtinFunction, found := builtinFunctions[wordLower]; found {
		builtinFunction(&forthStack)
	} else if definedWord, found := definedWords[word]; found {
		for _, subWord := range definedWord {
			executeWord(subWord)
		}
	} else {
		switch word {
		case "true":
			forthStack.Push(variant.ForthBool(true))
		case "false":
			forthStack.Push(variant.ForthBool(false))
		default:
			forthStack.Push(variant.ForthString(word))
		}
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	for {
		var input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var inQuotes = false
		var inputSplit = strings.FieldsFunc(input, func(r rune) bool {
			if r == '"' {
				inQuotes = !inQuotes
				return false
			}

			return !inQuotes && unicode.IsSpace(r)
		})

		if len(inputSplit) >= 4 && inputSplit[0] == ":" && inputSplit[len(inputSplit)-1] == ";" {
			definedWords[inputSplit[1]] = inputSplit[2 : len(inputSplit)-1]
		} else {
			for _, word := range inputSplit {
				executeWord(word)
			}
		}

		//fmt.Printf("%d: %v\n", forthStack.Size(), forthStack.Array())
	}
}
