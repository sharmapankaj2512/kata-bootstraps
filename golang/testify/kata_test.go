package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	ohce is a console application that echoes the reverse of what you input through the console.

    - [ ] "hello" -> "olleh"
    - [ ] "ohce" -> "echo"

	Even though it seems a silly application, ohce knows a thing or two.
		When you start oche, it greets you differently depending on the current time, but only in Spanish:
		Between 20 and 6 hours, ohce will greet you saying: ¡Buenas noches < your name >!
		Between 6 and 12 hours, ohce will greet you saying: ¡Buenos días < your name >!
		Between 12 and 20 hours, ohce will greet you saying: ¡Buenas tardes < your name >!

	When you introduce a palindrome, ohce likes it and after reverse-echoing it, it adds ¡Bonita palabra!

    ohce knows when to stop, you just have to write Stop! and it'll answer Adios < your name > and end.
 */
func TestFailing(t *testing.T) {
	assert.Equal(t, 42, doSomething("a"), "I'm failing you can start with me...")
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 42, doSomething("b"), "Answer to the Ultimate Question of Life, the Universe, and Everything")
}

func TestGen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "with a expected 0",
			input:    "a",
			expected: 0,
		},
		{
			name:     "with b expected 42",
			input:    "b",
			expected: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSomething(tt.input); got != tt.expected {
				t.Errorf("doSomething() = %v, but expected %v", got, tt.expected)
			}
		})
	}
}

func doSomething(input string) interface{} {
	if input == "b" {
		return 42
	}
	return 0
}
