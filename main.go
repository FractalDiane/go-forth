package main

import (
	"bufio"
	"log"
	"os"

	"goforth/forth"
)

func main() {
	var program = forth.NewForthProgram()

	switch len(os.Args) {
	case 1:
		var reader = bufio.NewReader(os.Stdin)
		for {
			var input, _ = reader.ReadString('\n')
			forth.ExecuteWordLine(&program, input)
		}
	case 2:
		if infile, err := os.Open(os.Args[1]); err == nil {
			var scanner = bufio.NewScanner(infile)
			for scanner.Scan() {
				forth.ExecuteWordLine(&program, scanner.Text())
			}
		} else {
			log.Fatalf("Error: Can't open file %s: %v", os.Args[1], err)
		}
	default:
		log.Fatalf("Error: Invalid argument count to go-forth (%d)", len(os.Args))
	}
}
