package main

import (
	"charrylang/repl"
	"fmt"
	"os"
	"os/user"
)

const logo string = `
_________ .__                                 .____                           
\_   ___ \|  |__ _____ ______________ ___.__. |    |   _____    ____    ____  
/    \  \/|  |  \\__  \\_  __ \_  __ <   |  | |    |   \__  \  /    \  / ___\ 
\     \___|   Y  \/ __ \|  | \/|  | \/\___  | |    |___ / __ \|   |  \/ /_/  >
 \______  /___|  (____  /__|   |__|   / ____| |_______ (____  /___|  /\___  / 
        \/     \/     \/              \/              \/    \/     \//_____/  
`

// open REPL
func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf(logo)
	fmt.Printf("Hello %s! This is the Charry programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command!\n")
	repl.Start(os.Stdin, os.Stdout)
}