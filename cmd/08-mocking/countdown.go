package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const finalWord = "GO!"
const countdownStart = 3

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}
