package main

import (
	"fmt"

	"github.com/dimensi0n/brainfuck"
)

func main() {
	const brain = "+.>++.>-."
	buffer := brainfuck.Compile(brain)
	fmt.Println(buffer)
}
