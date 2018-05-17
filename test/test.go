package main

import (
	"fmt"

	"github.com/dimensi0n/gobrainfuck"
)

func main() {
	const brain = ">+++.>-<--.>.>+"
	buffer := gobrainfuck.Compile(brain)
	fmt.Println(buffer)
}
