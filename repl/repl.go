package repl

import (
	"bufio"
	"fmt"
	"github.com/dollarkillerx/monkey/lexer"
	"github.com/dollarkillerx/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		text := scanner.Text()
		l := lexer.New(text)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
