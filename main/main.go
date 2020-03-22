package main

import (
	"fmt"

	"github.com/seemenkina/lrParser/parser"
)

func newTestGr() parser.Grammar {
	gr := parser.Grammar{}

	gr.AddRule("A", "!B!")
	gr.AddRule("B", "T+B")
	gr.AddRule("B", "T")
	gr.AddRule("T", "M")
	gr.AddRule("T", "M*T")
	gr.AddRule("M", "a")
	gr.AddRule("M", "b")
	gr.AddRule("M", "(B)")

	err := gr.AddNToken("A")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = gr.AddNToken("B")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = gr.AddNToken("T")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = gr.AddNToken("M")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	gr.AddTToken("!")
	gr.AddTToken("+")
	gr.AddTToken("*")
	gr.AddTToken("a")
	gr.AddTToken("b")
	gr.AddTToken("(")
	gr.AddTToken(")")

	gr.AddStartSymbol("A")
	return gr
}

func main() {
	gr := newTestGr()
	gr.PrintGrammar()

	lrParser := parser.LRParser{}
	lrParser.NewLRParser(gr, "!a*b!")

	err := lrParser.StartParse()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
