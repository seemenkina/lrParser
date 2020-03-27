package parser

import (
	"fmt"
	"testing"

	"github.com/seemenkina/lrParser/grammar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createGr1() grammar.Grammar {
	gr := grammar.Grammar{}

	gr.AddRule("B", "T+B")
	gr.AddRule("B", "T")
	gr.AddRule("T", "M")
	gr.AddRule("T", "M*T")
	gr.AddRule("M", "a")
	gr.AddRule("M", "b")
	gr.AddRule("M", "(B)")

	err := gr.AddNToken("B")
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

	gr.AddTToken("+")
	gr.AddTToken("*")
	gr.AddTToken("a")
	gr.AddTToken("b")

	gr.AddStartSymbol("B")
	return gr
}

func createGr2() grammar.Grammar {
	gr := grammar.Grammar{}

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

func TestLRParser_StartParse(t *testing.T) {

	tests := []struct {
		input  string
		gr     grammar.Grammar
		answer bool
	}{
		{"a+b", createGr1(), true},
		{"c+b", createGr1(), false},
		{"!(a+b)*(a+c)!", createGr2(), false},
		{"!(a+b)*(a+b)!", createGr2(), true},
		{"", createGr2(), false},
		{"", createGr1(), false},
		{"!a+b!", createGr2(), true},
	}

	t.Run("", func(t *testing.T) {
		for _, tt := range tests {
			var flag bool
			lrParser := LRParser{}
			lrParser.NewLRParser(tt.gr, tt.input)
			err := lrParser.StartParse()
			if err != nil {
				flag = false
			} else {
				flag = true
			}
			assert.EqualValues(t, flag, tt.answer)
		}
	})
}

func TestLRParser_PrintLRParserNOTOK(t *testing.T) {
	gr := createGr2()
	gr.PrintGrammar()
	input := "!a-b!"
	fmt.Printf("Input string: %s\n", input)
	lrParser := LRParser{}
	lrParser.NewLRParser(gr, input)
	err := lrParser.StartParse()
	fmt.Printf("Output: %s\n", err)
}

func TestLRParser_PrintLRParserOK(t *testing.T) {
	gr := createGr2()
	gr.PrintGrammar()
	input := "!a+b!"
	fmt.Printf("Input string: %s\n", input)
	lrParser := LRParser{}
	lrParser.NewLRParser(gr, input)
	err := lrParser.StartParse()
	require.NoError(t, err)
}
