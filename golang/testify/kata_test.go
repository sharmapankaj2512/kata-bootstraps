package kata

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
	ohce is a console application that echoes the reverse of what you input through the console.

    - [x] "hello" -> "olleh"
    - [x] "ohce" -> "echo"

	Even though it seems a silly application, ohce knows a thing or two.
		When you start oche, it greets you differently depending on the current time, but only in Spanish:
		Between 20 and 6 hours, ohce will greet you saying: ¡Buenas noches < your name >!
		- [x] Greet("Paul", 20) ->   ¡Buenas noches Paul!
		- [x] Greet("Paul", 5) ->   ¡Buenas noches Paul!
		Between 6 and 12 hours, ohce will greet you saying: ¡Buenos días < your name >!
		- [x] Greet("David", 6) ->   ¡Buenos días David!
		- [x] Greet("David", 11) ->   ¡Buenos días David!
		Between 12 and 20 hours, ohce will greet you saying: ¡Buenas tardes < your name >!
		- [ ] Greet("Daniel", 12) ->   ¡Buenas tardes Daniel!
		- [ ] Greet("Daniel", 19) ->   ¡Buenas tardes Daniel!
	When you introduce a palindrome, ohce likes it and after reverse-echoing it, it adds ¡Bonita palabra!

	[x] NewProgam with name "name" and time of 1AM, greet() -> ¡Buenos noches name!
	[ ] NewProgam with name "name" and time of 1AM, stop() -> ¡Adios name!

    ohce knows when to stop, you just have to write Stop! and it'll answer Adios < your name > and end.
 */

func Reverse(input string) string {
	var r = []rune(input)
	length := len(r)
	for i := 0; i < len(r) / 2; i++ {
		r[i], r[length-i-1] = r[length-i-1], r[i]
	}
	return string(r)
}

func TestHello(t *testing.T) {
	assert.Equal(t, "olleh", Reverse("hello"))
}

func TestEcho(t *testing.T) {
	assert.Equal(t, "ohce", Reverse("echo"))
}
func TestGreet(t *testing.T) {
	assert.Equal(t, "¡Buenas noches Paul!", Greet("Paul", 20))
	assert.Equal(t, "¡Buenas noches Paul!", Greet("Paul", 5))
}

func TestGreet_6_12(t *testing.T) {
	assert.Equal(t, "¡Buenos días David!", Greet("David", 6))
	assert.Equal(t, "¡Buenos días David!", Greet("David", 11))
}

func Greet(name string, hour int) string {
	if 6 <= hour && hour < 12 {
		return fmt.Sprintf("¡Buenos días %s!", name)
	}
	if 12 <= hour && hour < 20 {
		return "foo"
	}
	return fmt.Sprintf("¡Buenas noches %s!", name)
}

type Clock interface {
	now() time.Time
}

type RealClock struct{}

func (RealClock) now() time.Time {
	return time.Now()
}

type stubClock time.Time

func (t stubClock) now() time.Time {
	return time.Time(t)
}

type Program struct{
	Name string
	clock Clock
}

func (p *Program) greet() string {
	return Greet(p.Name, p.clock.now().Hour())
}

func (p *Program) depart() string {
	return fmt.Sprintf("¡Adios %s!", p.Name)
}

func (p *Program) Ohce(line string, w io.Writer) {
	rev := Reverse(line)
	fmt.Fprintln(w, rev)
	if rev == line {
		fmt.Println("bar")
	}
}

func NewProgram(name string) *Program {
	return &Program{
		Name: name,
		clock: RealClock{},
	}
}

const (
	nochesHour = 1
	diasHour = 11
	tardesHour = 17
)

func TestNewProgram(t *testing.T) {
	prog := NewProgram("name")
	prog.clock = stubClock(time.Date(2021, 02, 03, nochesHour, 0, 0, 0, time.Local))

	assert.Equal(t, prog.greet(), "¡Buenas noches name!")
}

func ohce(w io.Writer, prog *Program, r io.Reader) {
	fmt.Fprintln(w, prog.greet())

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "Stop!" {
			break
		}

		prog.Ohce(line, w)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, prog.depart())
}

func TestOhce(t *testing.T) {
 	input := strings.NewReader(`hello
echo
stop
Stop!
`)
	expected := []byte(`¡Buenos días Jorge!
olleh
ohce
pots
¡Adios Jorge!
`)

	prog := NewProgram("Jorge")
	prog.clock = stubClock(time.Date(2021, 02, 03, diasHour, 0, 0, 0, time.Local))

	output := new(bytes.Buffer)
	ohce(output, prog, input)

	assert.Equal(t, output.Bytes(), expected)
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
