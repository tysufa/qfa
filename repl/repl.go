package repl

import (
	"bufio"
	"fmt"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
	"io"
)

const PROMPT = "=> "

func Start(input io.Reader) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
