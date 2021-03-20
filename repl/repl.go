// repl/repl.go

package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeyc/lexer"
	"monkeyc/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		input := scanner.Text()

		theLexer := lexer.New(input)

		for tokenFound := theLexer.NextToken(); tokenFound.Type != token.EOF; tokenFound = theLexer.NextToken() {
			fmt.Printf("%+v\n", tokenFound)
		}
	}
}