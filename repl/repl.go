package repl

import (
	"github.com/tysufa/qfa/lexer"
)

func main() {
	l := lexer.New("")
	l.GetToken()
}
