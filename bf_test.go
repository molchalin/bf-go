package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	code := ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<+" +
		"+.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-" +
		"]<+."

	buf := new(bytes.Buffer)
	err := Interpret([]byte(code), bytes.NewReader(nil), buf)

	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", buf.String())
}
