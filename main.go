package main

import (
	"fmt"
	"github.com/tysufa/qfa/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the best programming language (it's the qfa so of course it's the best)\n", user.Username)
	fmt.Printf("WRITE\nCOMMANDS\nNOW\n")
	repl.Start(os.Stdin)
}
