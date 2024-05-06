package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')
		l := lexer.New(input)
		tk := l.GetToken()
		for tk.Type != token.EOF {
			fmt.Printf("%v : %v\n", tk.Type, tk.Value)
			tk = l.GetToken()
		}
		fmt.Print("\n")
	}
}
