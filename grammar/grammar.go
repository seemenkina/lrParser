package grammar

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Token type const
const (
	Term = iota
	NTerm
)

// Grammar's rules definition
type Rule struct {
	LSymbol string
	RSymbol string
}

// Non Terminal symbol
type NToken struct {
	N                string // symbol
	Alternative      []int  // alternative in rules
	CountAlternative int    // number of alternative in rules
}

// Terminal symbol
type TToken struct {
	t string
}

type Grammar struct {
	NTokens     []NToken
	tTokens     []TToken
	Rules       []Rule
	StartSymbol string
}

func (gr *Grammar) AddRule(ls, rs string) {
	rule := Rule{
		LSymbol: ls,
		RSymbol: rs,
	}
	gr.Rules = append(gr.Rules, rule)
}

func (gr *Grammar) AddTToken(ts string) {
	tt := TToken{
		t: ts,
	}
	gr.tTokens = append(gr.tTokens, tt)
}

func (gr *Grammar) AddNToken(ns string) error {
	var alt []int

	for i, r := range gr.Rules {
		if r.LSymbol == ns {
			alt = append(alt, i)
		}
	}
	if len(alt) == 0 {
		return fmt.Errorf("This symbol is not in the rules: %s ", ns)
	}

	nt := NToken{
		N:                ns,
		Alternative:      alt,
		CountAlternative: len(alt),
	}

	gr.NTokens = append(gr.NTokens, nt)
	return nil
}

func (gr *Grammar) AddStartSymbol(s string) {
	gr.StartSymbol = s
}

func (gr *Grammar) PrintGrammar() {
	fmt.Println("Rules:")
	for i, r := range gr.Rules {
		fmt.Printf("%d: %s -> %s\n", i, r.LSymbol, r.RSymbol)
	}
	fmt.Printf("Terminal Symbol: ")
	for _, t := range gr.tTokens {
		fmt.Printf("%s ", t.t)
	}
	fmt.Printf("\nStart Symbol: %s\n", gr.StartSymbol)
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

	_, _ = fmt.Fprintf(w, "\n%s\t%s\t%s\t", "Symbol", "Count Alt", "Alternative")
	_, _ = fmt.Fprintf(w, "\n%s\t%s\t%s\t", "------", "---------", "-----------")

	for _, n := range gr.NTokens {
		s := fmt.Sprintf("%d ", n.Alternative)
		_, _ = fmt.Fprintf(w, "\n%s\t%d\t%s\t", n.N, n.CountAlternative, s)
	}

}

func (gr *Grammar) FindNToken(s string) int {
	for i, nt := range gr.NTokens {
		if nt.N == s {
			return i
		}
	}
	return -1
}

func (gr *Grammar) IsNTerm(s string) int {
	if gr.FindNToken(s) == -1 {
		return Term
	}
	return NTerm
}
