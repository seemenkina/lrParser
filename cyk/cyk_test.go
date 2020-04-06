package cyk

import (
	"fmt"
	"testing"

	"github.com/seemenkina/lrParser/grammar"
)

func createGr1() grammar.Grammar {
	gr := grammar.Grammar{}

	gr.AddRule("S", "AA") // 0
	gr.AddRule("S", "AS") // 1
	gr.AddRule("S", "b")  // 2
	gr.AddRule("A", "SA") // 3
	gr.AddRule("A", "AS") // 4
	gr.AddRule("A", "a")  // 5

	gr.AddStartSymbol("S")
	return gr
}

func createGr2() grammar.Grammar {
	gr := grammar.Grammar{}

	gr.AddRule("S", "AB") // 0
	gr.AddRule("S", "BC") // 1
	gr.AddRule("A", "BA") // 2
	gr.AddRule("A", "a")  // 3
	gr.AddRule("B", "CC") // 4
	gr.AddRule("B", "b")  // 5
	gr.AddRule("C", "AB") // 6
	gr.AddRule("C", "a")  // 7

	gr.AddStartSymbol("S")
	return gr
}

func createGr3() grammar.Grammar {
	gr := grammar.Grammar{}

	gr.AddRule("B", "BC") // 0
	gr.AddRule("C", "PT") // 1
	gr.AddRule("P", "+")  // 2
	gr.AddRule("B", "TD") // 3
	gr.AddRule("D", "UM") // 4
	gr.AddRule("U", "*")  // 5
	gr.AddRule("B", "SE") // 6
	gr.AddRule("S", "(")  // 7
	gr.AddRule("E", "BR") // 8
	gr.AddRule("R", ")")  // 9
	gr.AddRule("B", "a")  // 10
	gr.AddRule("B", "b")  // 11
	gr.AddRule("B", "c")  // 12
	gr.AddRule("T", "TD") // 13
	gr.AddRule("T", "a")  // 14
	gr.AddRule("T", "b")  // 15
	gr.AddRule("T", "c")  // 16
	gr.AddRule("T", "SE") // 17
	gr.AddRule("M", "a")  // 18
	gr.AddRule("M", "b")  // 19
	gr.AddRule("M", "c")  // 20
	gr.AddRule("M", "SE") // 21

	gr.AddStartSymbol("B")
	return gr
}

func TestCYK(t *testing.T) {
	gr := createGr1()
	var cyk CYK
	cyk.NewCYK(gr, "babbb")
	ok := cyk.StartCYK()
	fmt.Println(ok)
	cyk.PrintOut()
}

func TestCYK2(t *testing.T) {
	gr := createGr2()
	var cyk CYK
	cyk.NewCYK(gr, "baaba")
	ok := cyk.StartCYK()
	fmt.Println(ok)
	cyk.PrintOut()
}

func TestCYK3(t *testing.T) {
	gr := createGr3()
	var cyk CYK
	cyk.NewCYK(gr, "(a+b)*d")
	ok := cyk.StartCYK()
	fmt.Println(ok)
	cyk.PrintOut()
}
