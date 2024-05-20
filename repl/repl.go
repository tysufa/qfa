package repl

import (
	"bufio"
	"fmt"
	"github.com/tysufa/qfa/token"
	"os"

	"github.com/tysufa/qfa/lexer"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')

		l := lexer.New(input)
		tok := l.GetToken()
		for tok.Type != token.EOF && tok.Type != token.NL {
			fmt.Printf("{%v : %v}\n", tok.Type, tok.Value)
			tok = l.GetToken()
		}
	}
}
