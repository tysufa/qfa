package repl

import (
	"fmt"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/tysufa/qfa/evaluator"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/parser"
)

const PROMPT = ">>> "

func printInput(input string, pos int) {
	cursor.ClearLine()   // on efface la ligne pour réécrire l'input modifié par dessus
	cursor.StartOfLine() // on se remet au début de la ligne pour pouvoir écrire depuis le début de la ligne sans espace bizarre
	fmt.Printf("%v", PROMPT+input)
	cursor.StartOfLine()
	cursor.Right(pos + len(PROMPT)) // on se déplace au dernier emplacement du curseur

}

func Run() {
	var input string = ""
	var inputs []string
	curInput := 0 // input actuel dans la liste des inputs déjà évalués
	curChar := 0  // position du charactère actuel (nécéssaire pour un affichage correct)
	printInput(input, curChar)

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		// TODO: tenter de faire un switch key.Code
		switch {
		case key.Code == keys.Up:
			if len(inputs) > 0 { // si on a déjà exécuté des commandes
				if curInput > 0 {
					curInput-- // on remonte à la commande précédente
				}
				// on récupére la string de la commande précédente et on place le curseur au dernier charactère
				input = inputs[curInput]
				curChar = len(input)
				printInput(input, curChar)

			}
		case key.Code == keys.Down:
			if curInput < len(inputs)-1 {
				curInput++
				input = inputs[curInput]
			} else if curInput == len(inputs)-1 {
				curInput++
				input = "" // si on est à la commande la plus récente, on remplace par une chaine vide pour simplifier l'execution de commandes
			} else {
				input = ""
			}
			curChar = len(input)
			printInput(input, curChar)

		case key.Code == keys.Left:
			// on déplace le curseur à gauche et on modifie donc la position du curseur curChar
			if curChar > 0 {
				curChar--
				cursor.Left(1)
			}

		case key.Code == keys.Right:
			if curChar < len(input) {
				curChar++
				cursor.Right(1)
			}

		// on quitte l'écoute des touches
		case key.Code == keys.CtrlC:
			return true, nil
		case key.Code == keys.Escape:
			return true, nil

		case key.Code == keys.Space:
			if curChar == len(input) {
				input += " "
			} else {
				input = input[:curChar] + " " + input[curChar:] //on insère l'espace à la position du curseur
			}
			curChar++
			printInput(input, curChar)

		case key.Code == keys.Backspace:
			if curChar >= 1 { // on vérifie qu'on ne tente pas de supprimer un charactère inexistant
				input = input[:curChar-1] + input[curChar:]
				curChar--
			}
			printInput(input, curChar)

		case key.Code == keys.Enter:
			l := lexer.New(input)
			// tok := l.GetToken()
			// for tok.Type != token.EOF && tok.Type != token.NL {
			// 	fmt.Printf("{%v : %v}\n", tok.Type, tok.Value)
			// 	tok = l.GetToken()
			// }
			p := parser.New(l)
			stmts := p.GetStatements()
			evaluated := evaluator.EvaluateStatements(stmts)
			for _, ev := range evaluated {
				if ev != nil {
					fmt.Printf("\n%v\n", ev.Inspect())
				} else {
					fmt.Printf("\n\n")
				}
			}

			// fmt.Printf("\nEvaluate(%v)\n", input)

			inputs = append(inputs, input) // on ajoute la commande à la liste de commandes exécutés

			// on reset la position du curseur, incrémente le nombre de commandes exécutés et reset l'input
			curChar = 0
			curInput = len(inputs) - 1
			curInput++
			input = ""
			printInput(input, curChar)

		case key.Code == keys.CtrlU: // effacer toute la ligne
			input = ""
			curChar = 0
			printInput(input, curChar)

		default:
			if len(key.String()) == 1 { //on évite de détecter les chaines types ctrlA, esc, up, ... en vérifiant qu'il n'y a qu'un seul charactère
				if curChar == len(input) {
					input += key.String()
				} else {
					input = input[:curChar] + key.String() + input[curChar:]
				}
				curChar++
				printInput(input, curChar)

			}
		}

		return false, nil
	})
}
