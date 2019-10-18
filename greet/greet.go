package greet

import (
	"fmt"
	"io"
	"os"
)

// Greet accepts a writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
