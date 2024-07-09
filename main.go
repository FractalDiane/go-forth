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

type forthProgram struct {
	forthStack   stack.Stack[variant.Variant]
	definedWords map[string][]string

	wordIndex uint
	loopStack stack.Stack[uint]
}

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

func printTop(program *forthProgram) {
	var top = program.forthStack.Top()
	fmt.Printf("%v", *top)
	program.forthStack.Pop()
}

func printTopLn(program *forthProgram) {
	var top = program.forthStack.Top()
	fmt.Printf("%v\n", *top)
	program.forthStack.Pop()
}

func emitTop(program *forthProgram) {
	var top = program.forthStack.Top()
	switch topCast := (*top).(type) {
	case variant.ForthInt:
		fmt.Printf("%c", rune(topCast))
		program.forthStack.Pop()
	default:
		panic("test")
	}
}

func drop(program *forthProgram) {
	program.forthStack.Pop()
}

func dup(program *forthProgram) {
	var top = *program.forthStack.Top()
	program.forthStack.Push(top)
}

func swap(program *forthProgram) {
	program.forthStack.SwapTopElements()
}

func over(program *forthProgram) {
	var second = *program.forthStack.Second()
	program.forthStack.Push(second)
}

func rotate(program *forthProgram) {
	program.forthStack.RotateTopElements()
}

func random(program *forthProgram) {
	var value = rand.Int64()
	program.forthStack.Push(variant.ForthInt(value))
}

func randomf(program *forthProgram) {
	var value = rand.Float64()
	program.forthStack.Push(variant.ForthFloat(value))
}

func beginLoop(program *forthProgram) {
	program.loopStack.Push(program.wordIndex)
}

func loopAgain(program *forthProgram) {
	program.wordIndex = *program.loopStack.Top()
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

var builtinFunctions = map[string]func(*forthProgram){
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

	"begin": beginLoop,
	"again": loopAgain,
}

func executeWord(program *forthProgram, word string) {
	var wordLower = strings.ToLower(word)
	if integer, err := strconv.Atoi(word); err == nil {
		program.forthStack.Push(variant.ForthInt(integer))
	} else if float, err := strconv.ParseFloat(word, 64); err == nil {
		program.forthStack.Push(variant.ForthFloat(float))
	} else if strings.HasPrefix(word, `"`) && strings.HasSuffix(word, `"`) {
		var str = strings.TrimPrefix(word, `"`)
		str = strings.TrimSuffix(str, `"`)
		program.forthStack.Push(variant.ForthString(str))
	} else if binOpFunction, found := binaryOperators[wordLower]; found {
		var rhs = *program.forthStack.Top()
		program.forthStack.Pop()
		var lhs = *program.forthStack.Top()
		program.forthStack.Pop()
		program.forthStack.Push(binOpFunction(lhs, rhs))
	} else if unOpFunction, found := unaryOperators[wordLower]; found {
		var operand = *program.forthStack.Top()
		program.forthStack.Pop()
		program.forthStack.Push(unOpFunction(operand))
	} else if builtinFunction, found := builtinFunctions[wordLower]; found {
		builtinFunction(program)
	} else if definedWord, found := program.definedWords[word]; found {
		for _, subWord := range definedWord {
			executeWord(program, subWord)
		}
	} else {
		switch word {
		case "true":
			program.forthStack.Push(variant.ForthBool(true))
		case "false":
			program.forthStack.Push(variant.ForthBool(false))
		default:
			program.forthStack.Push(variant.ForthString(word))
		}
	}
}

func executeWordLine(program *forthProgram, wordLine string) {
	wordLine = strings.TrimSpace(wordLine)

	var inQuotes = false
	var inputSplit = strings.FieldsFunc(wordLine, func(r rune) bool {
		if r == '"' {
			inQuotes = !inQuotes
			return false
		}

		return !inQuotes && unicode.IsSpace(r)
	})

	if len(inputSplit) >= 4 && inputSplit[0] == ":" && inputSplit[len(inputSplit)-1] == ";" {
		program.definedWords[inputSplit[1]] = inputSplit[2 : len(inputSplit)-1]
	} else {
		for program.wordIndex < uint(len(inputSplit)) {
			var thisWord = inputSplit[program.wordIndex]
			executeWord(program, thisWord)
			program.wordIndex++
		}
	}
}

func main() {
	var program forthProgram
	program.definedWords = make(map[string][]string, 5)

	switch len(os.Args) {
	case 1:
		var reader = bufio.NewReader(os.Stdin)
		for {
			var input, _ = reader.ReadString('\n')
			executeWordLine(&program, input)
		}
	case 2:
		if infile, err := os.Open(os.Args[1]); err == nil {
			var scanner = bufio.NewScanner(infile)
			for scanner.Scan() {
				executeWordLine(&program, scanner.Text())
			}
		} else {
			panic("test")
		}
	default:
		panic("test")
	}
}
