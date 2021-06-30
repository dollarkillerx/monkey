package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func trace(msg string) string {
	//incIdent()
	//tracePrint("BEGIN " + msg)
	return msg
}

func unTrace(msg string) {
	//tracePrint("END " + msg)
	//decIdent()
}
