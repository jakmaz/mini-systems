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
	fmt.Println("+-------------------------------------------------------------------+")
	fmt.Printf("Hello %s! This is the minilang programming language!\n", user.Username)
	fmt.Println("Type in your code and press enter to run it.")
	fmt.Println("Use /tokens, /parse, /eval, /compile and /run to switch between modes")
	fmt.Println("Use /help for help and /exit to quit")
	fmt.Println("+-------------------------------------------------------------------+")
	repl.Start(os.Stdin, os.Stdout)
}
