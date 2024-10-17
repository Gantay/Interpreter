package main

import (
	"fmt"
	"github.com/Gantay/Interpreter/repl"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic("Who are you?")
	}

	fmt.Printf("Hello %s! This is Monkey programming language!\n", user.Username)

	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
