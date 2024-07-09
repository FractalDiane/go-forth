package tests

import (
	"fmt"
	"goforth/forth"
	"goforth/variant"
	"testing"
)

func runTestLine(line string, expectedValues ...variant.Variant) (passed bool, err string) {
	var program = forth.NewForthProgram()
	forth.ExecuteWordLine(&program, line)

	for _, expectedValue := range expectedValues {
		if program.StackTop() == nil {
			if expectedValue == nil {
				continue
			} else {
				return false, fmt.Sprintf("\nExpression: %v\nExpected: %v\nGot: nil", line, expectedValue)
			}
		}

		var actualValue = *program.StackTop()
		if actualValue != expectedValue {
			return false, fmt.Sprintf("\nExpression: %v\nExpected: %v\nGot: %v", line, expectedValue, actualValue)
		}

		program.StackPop()
	}

	return true, ""
}

func TestAdd(t *testing.T) {
	if passed, err := runTestLine("5 3 +", variant.ForthInt(8)); !passed {
		t.Fatal(err)
	}
}

func TestSubtract(t *testing.T) {
	if passed, err := runTestLine("5 3 -", variant.ForthInt(2)); !passed {
		t.Fatal(err)
	}
}

func TestMultiply(t *testing.T) {
	if passed, err := runTestLine("5 3 *", variant.ForthInt(15)); !passed {
		t.Fatal(err)
	}
}

func TestDivide(t *testing.T) {
	if passed, err := runTestLine("5 3 /", variant.ForthInt(1)); !passed {
		t.Fatal(err)
	}
}

func TestModulus(t *testing.T) {
	if passed, err := runTestLine("5 3 %", variant.ForthInt(2)); !passed {
		t.Fatal(err)
	}
}

func TestDuplicate(t *testing.T) {
	if passed, err := runTestLine("37 dup", variant.ForthInt(37), variant.ForthInt(37)); !passed {
		t.Fatal(err)
	}
}

func TestSwap(t *testing.T) {
	if passed, err := runTestLine("123 19 swap", variant.ForthInt(123), variant.ForthInt(19)); !passed {
		t.Fatal(err)
	}
}

func TestDrop(t *testing.T) {
	if passed, err := runTestLine("42 12 drop", variant.ForthInt(42), nil); !passed {
		t.Fatal(err)
	}
}

func TestOver(t *testing.T) {
	if passed, err := runTestLine("78 6 over", variant.ForthInt(78), variant.ForthInt(6), variant.ForthInt(78)); !passed {
		t.Fatal(err)
	}
}

func TestRot(t *testing.T) {
	if passed, err := runTestLine("10 20 30 rot", variant.ForthInt(10), variant.ForthInt(30), variant.ForthInt(20)); !passed {
		t.Fatal(err)
	}
}

func TestIfTrue(t *testing.T) {
	if passed, err := runTestLine(`1 if "true" else "false" then`, variant.ForthString("true")); !passed {
		t.Fatal(err)
	}
}

func TestIfFalse(t *testing.T) {
	if passed, err := runTestLine(`0 if "true" else "false" then`, variant.ForthString("false")); !passed {
		t.Fatal(err)
	}
}

func TestDoLoop(t *testing.T) {
	if passed, err := runTestLine("5 0 do i loop", variant.ForthInt(4), variant.ForthInt(3), variant.ForthInt(2), variant.ForthInt(1), variant.ForthInt(0)); !passed {
		t.Fatal(err)
	}
}

func TestUntilLoop(t *testing.T) {
	if passed, err := runTestLine("5 begin 1 - dup until", variant.ForthInt(0), nil); !passed {
		t.Fatal(err)
	}
}
