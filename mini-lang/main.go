package main

import (
	"fmt"
	"minilang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the minilang programming language!\n", user.Username)
	fmt.Printf("Type in your code and press enter to run it.\n")
	repl.Start(os.Stdin, os.Stdout)
}
