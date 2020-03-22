package lrparser

import (
	"fmt"
	"strings"
)

type L1Token struct {
	token     string
	tokenType int
	// for non terminal symbol
	countAlternative int // number of alternative in rules
	firstN           int // first alternative in rules
}

// Token type const
const (
	Term = iota
	NTerm
)

type L2Token struct {
	token     string
	tokenType int
}

// State const
const (
	normal = iota
	ret
	end
)

type LRParser struct {
	grammar Grammar
	input   string
	l1Stack []L1Token
	l2Stack []L2Token
	state   int
	output  string
}

func (lrp *LRParser) newLRParser(gr Grammar, in string) {
	lrp.grammar = gr
	lrp.input = in
	lrp.state = normal
	for _, nt := range lrp.grammar.nTokens {
		if nt.n == lrp.grammar.startSymbol {
			lrp.l2Stack = append(lrp.l2Stack, L2Token{
				token:     nt.n,
				tokenType: NTerm})
		}
	}
}

// (q, i, α, Aβ ) |- (q, i, αA1, γ1β )
// A -> γ1
// A1 - first alternative for A
func (lrp *LRParser) growthOfTree(k int) int {

	return 0
}

func (lrp *LRParser) successEnd() {

}

func (lrp *LRParser) startParse() error {
	i := 0
	j := 0
	k := 1
	for {
		switch lrp.state {
		case normal:
			switch {
			case lrp.l2Stack[k-1].tokenType == NTerm:
				// f1
				continue

			case lrp.l2Stack[k-1].tokenType == Term &&
				strings.Compare(lrp.l2Stack[k-1].token, string(lrp.input[i])) != 0:
				// f4
				continue

			case lrp.l2Stack[k-1].tokenType == Term &&
				strings.Compare(lrp.l2Stack[k-1].token, string(lrp.input[i])) == 0 && i == len(lrp.input):
				switch len(lrp.l2Stack) {
				case 0:
					// f3
					continue
				default:
					// f3'
					continue
				}

			case lrp.l2Stack[k-1].tokenType == Term &&
				strings.Compare(lrp.l2Stack[k-1].token, string(lrp.input[i])) == 0 && i != len(lrp.input):
				switch len(lrp.l2Stack) {
				case 0:
					// f3 '
					continue
				default:
					continue
				}
			}

		case ret:
			switch {
			case lrp.l1Stack[j-1].tokenType == Term:
				// f5
				continue
			case lrp.l1Stack[j-1].tokenType == NTerm && lrp.l1Stack[j-1].countAlternative != 0:
				// f6a
				continue
			case lrp.l1Stack[j-1].tokenType == NTerm && lrp.l1Stack[j-1].countAlternative == 0:
				if lrp.l1Stack[j-1].token == lrp.grammar.startSymbol && i == 0 {
					// f6 b
					return fmt.Errorf("The input string does not belong to the grammar ")
				} else {
					// f6 v
					continue
				}
			}

		case end:
			lrp.successEnd()
			return nil
		}
	}
}
