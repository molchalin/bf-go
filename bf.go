package main

import (
	"fmt"
	"io"
)

func Interpret(code []byte, in io.Reader, out io.Writer) error {
	jumpStack := []int{}

	jumpTo := make(map[int]int)

	for i, c := range code {
		switch c {
		case '[':
			jumpStack = append(jumpStack, i)
		case ']':
			if len(jumpStack) == 0 {
				return fmt.Errorf("cycle error. pos: %v", i)
			}
			to := jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]
			jumpTo[to] = i
			jumpTo[i] = to
		}
	}

	var ptr int
	data := make([]byte, 1)

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++
			if ptr == len(data) {
				data = append(data, 0)
			}
		case '<':
			if ptr == 0 {
				return fmt.Errorf("can't decrement data pointer. pos: %v", i)
			}
			ptr--
		case '+':
			data[ptr]++
		case '-':
			data[ptr]--
		case '.':
			_, err := out.Write([]byte{data[ptr]})
			if err != nil {
				return fmt.Errorf("can't print char: %v", err)
			}
		case ',':
			_, err := in.Read(data[ptr : ptr+1])
			if err != nil {
				return fmt.Errorf("can't read char: %v", err)
			}
		case '[':
			if data[ptr] == 0 {
				i = jumpTo[i]
			}
		case ']':
			if data[ptr] != 0 {
				i = jumpTo[i]
			}
		}
	}
	return nil
}
