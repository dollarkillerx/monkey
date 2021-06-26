package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/dollarkillerx/monkey/repl"
)

func main() {
	current, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", current.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
