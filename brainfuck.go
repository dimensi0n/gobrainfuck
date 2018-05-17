// Package gobrainfuck is the unofficial brainfuck package
package gobrainfuck

import (
	"strings"
)

// Compile function allow you to compile a string to brainfuck
func Compile(b string) []int {
	s := strings.SplitAfter(b, "")
	var buffer = make([]int, len(s))
	pointer := 0

	for i := range s {
		switch s[i] {
		case ">":
			pointer++
		case "<":
			pointer--
		case "+":
			buffer[pointer]++
		case "-":
			buffer[pointer]--
		case ".":
			println(buffer[pointer])
		}
	}

	return buffer
}
