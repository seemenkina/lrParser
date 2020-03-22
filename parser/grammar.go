package parser

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Rule struct {
	lSymbol string
	rSymbol string
}

// Non Terminal symbol
type NToken struct {
	n                string // symbol
	alternative      []int  // alternative in rules
	countAlternative int    // number of alternative in rules
	firstN           int    // first alternative in rules
}

// Terminal symbol
type TToken struct {
	t string
}

type Grammar struct {
	nTokens     []NToken
	tTokens     []TToken
	rules       []Rule
	startSymbol string
}

func (gr *Grammar) AddRule(ls, rs string) {
	rule := Rule{
		lSymbol: ls,
		rSymbol: rs,
	}
	gr.rules = append(gr.rules, rule)
}

func (gr *Grammar) AddTToken(ts string) {
	tt := TToken{
		t: ts,
	}
	gr.tTokens = append(gr.tTokens, tt)
}

func (gr *Grammar) AddNToken(ns string) error {
	var alt []int

	for i, r := range gr.rules {
		if r.lSymbol == ns {
			alt = append(alt, i)
		}
	}
	if len(alt) == 0 {
		return fmt.Errorf("This symbol is not in the rules: %s ", ns)
	}

	nt := NToken{
		n:                ns,
		alternative:      alt,
		countAlternative: len(alt),
		firstN:           alt[0],
	}

	gr.nTokens = append(gr.nTokens, nt)
	return nil
}

func (gr *Grammar) AddStartSymbol(s string) {
	gr.startSymbol = s
}

func (gr *Grammar) PrintGrammar() {
	fmt.Println("Rules:")
	for i, r := range gr.rules {
		fmt.Printf("%d: %s -> %s\n", i, r.lSymbol, r.rSymbol)
	}
	fmt.Printf("Terminal Symbol: ")
	for _, t := range gr.tTokens {
		fmt.Printf("%s ", t.t)
	}
	fmt.Printf("\nStart Symbol: %s\n", gr.startSymbol)
	fmt.Print("Non Terminal:")
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer func() {
		_, _ = fmt.Fprintf(w, "\n")
		err := w.Flush()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}()

	_, _ = fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t", "Symbol", "Count Alt", "First ALter", "Alternative")
	_, _ = fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t", "------", "---------", "-----------", "-----------")

	for _, n := range gr.nTokens {
		s := fmt.Sprintf("%d ", n.alternative)
		_, _ = fmt.Fprintf(w, "\n%s\t%d\t%d\t%s\t", n.n, n.countAlternative, n.firstN, s)
	}

}

func (gr *Grammar) findNToken(s string) int {
	for i, nt := range gr.nTokens {
		if nt.n == s {
			return i
		}
	}
	return -1
}

func (gr *Grammar) findTToken(s string) int {
	for i, tt := range gr.tTokens {
		if tt.t == s {
			return i
		}
	}
	return -1
}

func (gr *Grammar) IsNTerm(s string) int {
	if gr.findNToken(s) == -1 {
		return Term
	}
	return NTerm
}
