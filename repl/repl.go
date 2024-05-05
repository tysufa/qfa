package repl

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/tysufa/qfa/lexer"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	}
}
