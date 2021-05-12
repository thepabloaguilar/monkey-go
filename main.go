package main

import (
	"fmt"
	"monkey-go/repl"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", currentUser.Username)
	fmt.Println("Feel free to type in commands")

	repl.Start(os.Stdin, os.Stdout)
}
