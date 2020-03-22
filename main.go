package lrparser

func newTestGr() Grammar {
	gr := Grammar{}

	gr.addRule("A", "!B!")
	gr.addRule("B", "T+B")
	gr.addRule("B", "T")
	gr.addRule("T", "M")
	gr.addRule("T", "M*T")
	gr.addRule("M", "a")
	gr.addRule("M", "b")
	gr.addRule("M", "(B)")

	err := gr.addNToken("A")
	if err != nil {
		panic(err)
	}

	err = gr.addNToken("B")
	if err != nil {
		panic(err)
	}

	err = gr.addNToken("T")
	if err != nil {
		panic(err)
	}

	err = gr.addNToken("M")
	if err != nil {
		panic(err)
	}

	gr.addTToken("!")
	gr.addTToken("+")
	gr.addTToken("*")
	gr.addTToken("a")
	gr.addTToken("b")
	gr.addTToken("(")
	gr.addTToken(")")
	gr.addStartSymbol("A")
	return gr
}

func main() {
	gr := newTestGr()
	gr.printGrammar()

	lrParser := LRParser{}
	lrParser.newLRParser(gr, "!a+b!")

}
