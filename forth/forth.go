package forth

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
	"strings"
	"unicode"

	"goforth/stack"
	"goforth/variant"
)

type loopEntry struct {
	isDoLoop     bool
	loopIndex    int
	lowerBound   variant.ForthInt
	upperBound   variant.ForthInt
	currentValue variant.ForthInt
}

type ForthProgram struct {
	forthStack   stack.Stack[variant.Variant]
	definedWords map[string][]string

	wordIndex int
	loopStack stack.Stack[loopEntry]
}

func NewForthProgram() ForthProgram {
	var program ForthProgram
	program.definedWords = make(map[string][]string, 5)
	return program
}

func (program *ForthProgram) StackTop() *variant.Variant {
	return program.forthStack.Top()
}

func (program *ForthProgram) StackPop() {
	program.forthStack.Pop()
}

func (program *ForthProgram) Reset() {
	program.forthStack.Clear()
	program.loopStack.Clear()
	program.definedWords = make(map[string][]string, 5)
	program.wordIndex = 0
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

func printTop(program *ForthProgram) {
	if !program.forthStack.IsEmpty() {
		var top = program.forthStack.Top()
		fmt.Printf("%v", *top)
		program.forthStack.Pop()
	} else {
		log.Fatal("Error: Attempted to print, but the stack is empty")
	}
}

func printTopLn(program *ForthProgram) {
	if !program.forthStack.IsEmpty() {
		var top = program.forthStack.Top()
		fmt.Printf("%v\n", *top)
		program.forthStack.Pop()
	} else {
		log.Fatal("Error: Attempted to print, but the stack is empty")
	}
}

func emitTop(program *ForthProgram) {
	if !program.forthStack.IsEmpty() {
		var top = program.forthStack.Top()
		switch topCast := (*top).(type) {
		case variant.ForthInt:
			fmt.Printf("%c", rune(topCast))
			program.forthStack.Pop()
		default:
			log.Fatal("Error: emit failed to convert its argument")
		}
	} else {
		log.Fatal("Error: Attempted to emit, but the stack is empty")
	}
}

func drop(program *ForthProgram) {
	program.forthStack.Pop()
}

func dup(program *ForthProgram) {
	var top = *program.forthStack.Top()
	program.forthStack.Push(top)
}

func swap(program *ForthProgram) {
	program.forthStack.SwapTopElements()
}

func over(program *ForthProgram) {
	var second = *program.forthStack.Second()
	program.forthStack.Push(second)
}

func rotate(program *ForthProgram) {
	program.forthStack.RotateTopElements()
}

func random(program *ForthProgram) {
	var value = rand.Int64()
	program.forthStack.Push(variant.ForthInt(value))
}

func randomf(program *ForthProgram) {
	var value = rand.Float64()
	program.forthStack.Push(variant.ForthFloat(value))
}

func beginLoop(program *ForthProgram) {
	program.loopStack.Push(loopEntry{false, program.wordIndex, 0, 0, 0})
}

func loopAgain(program *ForthProgram) {
	var topEntry = program.loopStack.Top()
	if topEntry != nil && !topEntry.isDoLoop {
		program.wordIndex = topEntry.loopIndex
	} else {
		log.Fatalf("Error: Mismatched 'again'; no matching 'begin'")
	}
}

func loopUntil(program *ForthProgram) {
	var topEntry = program.loopStack.Top()
	if topEntry != nil && !topEntry.isDoLoop {
		var flag = *program.StackTop()
		program.StackPop()
		if flag.AsBool() {
			program.wordIndex = topEntry.loopIndex
		}
	} else {
		log.Fatalf("Error: Mismatched 'until'; no matching 'begin'")
	}
}

func doLoopStart(program *ForthProgram) {
	var lowerBound = (*program.StackTop()).(variant.ForthInt)
	program.StackPop()
	var upperBound = (*program.StackTop()).(variant.ForthInt)
	program.StackPop()
	program.loopStack.Push(loopEntry{true, program.wordIndex, lowerBound, upperBound, lowerBound})
}

func doLoopLoop(program *ForthProgram) {
	var topEntry = program.loopStack.Top()
	if topEntry != nil && topEntry.isDoLoop {
		if topEntry.currentValue < topEntry.upperBound-1 {
			program.wordIndex = topEntry.loopIndex
			topEntry.currentValue++
		}
	} else {
		log.Fatalf("Error: Mismatched 'loop'; no matching 'do'")
	}
}

func loopIndex(program *ForthProgram) {
	var topEntry = program.loopStack.Top()
	if topEntry != nil && topEntry.isDoLoop {
		program.forthStack.Push(variant.ForthInt(topEntry.currentValue))
	} else {
		log.Fatalf("Error: 'i' has no corresponding loop to query")
	}
}

func loopIndex2(program *ForthProgram) {
	var topEntry = program.loopStack.Peek(1)
	if topEntry != nil && topEntry.isDoLoop {
		program.forthStack.Push(variant.ForthInt(topEntry.currentValue))
	} else {
		log.Fatalf("Error: 'i' has no corresponding loop to query")
	}
}

func loopIndex3(program *ForthProgram) {
	var topEntry = program.loopStack.Peek(2)
	if topEntry != nil && topEntry.isDoLoop {
		program.forthStack.Push(variant.ForthInt(topEntry.currentValue))
	} else {
		log.Fatalf("Error: 'i' has no corresponding loop to query")
	}
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

var builtinFunctions = map[string]func(*ForthProgram){
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
	"until": loopUntil,
	"do":    doLoopStart,
	"loop":  doLoopLoop,
	"i":     loopIndex,
	"j":     loopIndex2,
	"k":     loopIndex3,
}

func ExecuteWord(program *ForthProgram, word string) {
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
			ExecuteWord(program, subWord)
		}
	} else {
		switch word {
		case "true":
			program.forthStack.Push(variant.ForthBool(true))
		case "false":
			program.forthStack.Push(variant.ForthBool(false))
		default:
			log.Fatalf("Error: Unrecognized word '%s'", word)
		}
	}
}

func ExecuteWordLine(program *ForthProgram, wordLine string) {
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
		program.wordIndex = 0
		for program.wordIndex < len(inputSplit) {
			var thisWord = inputSplit[program.wordIndex]
			ExecuteWord(program, thisWord)
			program.wordIndex++
		}
	}
}
