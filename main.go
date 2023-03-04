package main

import (
	"charrylang/repl"
	"fmt"
	"os"
	"os/user"
)

// open REPL
func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Charry programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command!\n")
	repl.Start(os.Stdin, os.Stdout)
}