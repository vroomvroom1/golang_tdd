package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper interface signifies that the sleep method is required
type Sleeper interface {
	Sleep()
}

// SpySleeper represents a struct holding the amount of calls
type SpySleeper struct {
	Calls int
}

// Sleep function adds one count to the calls value
func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown takes in a variable of type buffer and returns it
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
