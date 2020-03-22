package parser

import (
	"fmt"
)

type L1Token struct {
	token     string
	tokenType int
	// for non terminal symbol
	countAlternative int // number of alternative in rules
	numOfAlternative int // current number of alternative in rules
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
	output  []int
	l1Iter  int
	l2Iter  int
	inIter  int
}

func (lrp *LRParser) NewLRParser(gr Grammar, in string) {
	lrp.grammar = gr
	lrp.input = in
	lrp.state = normal
	for _, nt := range lrp.grammar.nTokens {
		if nt.n == lrp.grammar.startSymbol {
			lrp.l2Stack = append(lrp.l2Stack, L2Token{
				token:     nt.n,
				tokenType: NTerm,
			})
		}
	}
}

func (lrp *LRParser) PushL1Stack(l1t L1Token) {
	newL1Stack := make([]L1Token, 1)
	newL1Stack[0] = l1t
	newL1Stack = append(newL1Stack, lrp.l1Stack...)
	lrp.l1Stack = newL1Stack
}

func (lrp *LRParser) PushL2Stack(l2t L2Token) {
	newL2Stack := make([]L2Token, len(l2t.token))
	for i, s := range l2t.token {
		newL2Stack[i].token = string(s)
		newL2Stack[i].tokenType = lrp.grammar.IsNTerm(string(s))
	}

	newL2Stack = append(newL2Stack, lrp.l2Stack...)
	lrp.l2Stack = newL2Stack
}

// (q, i, α, Aβ ) |- (q, i, αA1, γ1β )
// A -> γ1
// A1 - first alternative for A
func (lrp *LRParser) growthOfTree() {
	sym := lrp.l2Stack[lrp.l2Iter].token
	nTerm := lrp.grammar.nTokens[lrp.grammar.findNToken(sym)]
	l1t := L1Token{
		token:            nTerm.n,
		tokenType:        NTerm,
		countAlternative: nTerm.countAlternative,
		numOfAlternative: 1,
	}

	lrp.PushL1Stack(l1t)

	numRule := nTerm.alternative[l1t.numOfAlternative-1]
	rRule := lrp.grammar.rules[numRule].rSymbol
	l2t := L2Token{
		token:     rRule,
		tokenType: lrp.grammar.IsNTerm(rRule),
	}

	lrp.l2Stack = lrp.l2Stack[1:]
	lrp.PushL2Stack(l2t)
	return
}

// (q, i, α, aβ ) |- (q, i+1, αa, β )
// a = ai, i  ≤ n
func (lrp *LRParser) successfulCompareInputCharacter() {
	lrp.inIter++
	l1t := L1Token{
		token:            lrp.l2Stack[0].token,
		tokenType:        Term,
		countAlternative: 0,
		numOfAlternative: 1,
	}

	lrp.l2Stack = lrp.l2Stack[1:]
	lrp.PushL1Stack(l1t)
}

//  (q, n + 1, α, e ) |- ( t, n + 1, α, e)
func (lrp *LRParser) successfulCompletion() {
	lrp.state = end

	for _, l1t := range lrp.l1Stack {
		if l1t.tokenType == Term {
			continue
		}
		it := lrp.grammar.findNToken(l1t.token)
		r := lrp.grammar.nTokens[it].alternative[l1t.numOfAlternative-1]
		lrp.output = append(lrp.output, r)
	}

	for i := len(lrp.output)/2 - 1; i >= 0; i-- {
		opp := len(lrp.output) - 1 - i
		lrp.output[i], lrp.output[opp] = lrp.output[opp], lrp.output[i]
	}
}

// (b, i, αa, β ) |- (b, i – 1, α, aβ )
func (lrp *LRParser) returnOnTerm() {
	lrp.inIter--
	l2t := L2Token{
		token:     lrp.l1Stack[lrp.l1Iter].token,
		tokenType: Term,
	}
	lrp.l1Stack = lrp.l1Stack[1:]
	lrp.PushL2Stack(l2t)
}

func (lrp *LRParser) testAlternative() {
	lrp.state = normal

	lrp.l1Stack[0].numOfAlternative++

	it := lrp.grammar.findNToken(lrp.l1Stack[0].token)
	r := lrp.grammar.nTokens[it].alternative[lrp.l1Stack[0].numOfAlternative-1]

	rRule := lrp.grammar.rules[r].rSymbol
	orRule := lrp.grammar.rules[r-1].rSymbol
	lrp.l2Stack = lrp.l2Stack[len(orRule):]
	lrp.PushL2Stack(L2Token{
		token:     rRule,
		tokenType: -1,
	})

	// lrp.l2Stack[0].token = rRule
	// lrp.l2Stack[0].tokenType = lrp.grammar.IsNTerm(rRule)

}

func (lrp *LRParser) returnNonTerm() {
	it := lrp.grammar.findNToken(lrp.l1Stack[0].token)
	nr := lrp.grammar.nTokens[it].alternative[lrp.l1Stack[0].numOfAlternative-1]
	rRule := lrp.grammar.rules[nr].rSymbol
	lRule := lrp.grammar.rules[nr].lSymbol
	lrp.l2Stack = lrp.l2Stack[len(rRule):]
	lrp.PushL2Stack(L2Token{
		token:     lRule,
		tokenType: -1,
	})

	// lrp.l2Stack[0].token = lRule
	// lrp.l2Stack[0].tokenType = lrp.grammar.IsNTerm(lRule)

	lrp.l1Stack = lrp.l1Stack[1:]
}

func (lrp *LRParser) printOutput() {
	fmt.Printf("Left Out: ")
	for _, o := range lrp.output {
		fmt.Printf("%d ", o)
	}
	fmt.Println()
}

func (lrp *LRParser) StartParse() error {
	for {
		switch lrp.state {
		case normal:
			switch {
			case lrp.l2Stack[lrp.l2Iter].tokenType == NTerm:
				lrp.growthOfTree()
				continue

			case lrp.l2Stack[lrp.l2Iter].tokenType == Term &&
				lrp.l2Stack[lrp.l2Iter].token != string(lrp.input[lrp.inIter]):
				lrp.state = ret
				continue

			case lrp.l2Stack[lrp.l2Iter].tokenType == Term && lrp.l2Stack[lrp.l2Iter].token == string(lrp.input[lrp.inIter]):
				lrp.successfulCompareInputCharacter()
				if lrp.inIter == len(lrp.input) {
					switch len(lrp.l2Stack) {
					case 0:
						lrp.successfulCompletion()
						continue
					default:
						lrp.state = ret
						continue
					}
				} else {
					switch len(lrp.l2Stack) {
					case 0:
						lrp.state = ret
						continue
					default:
						continue
					}
				}

			}

		case ret:
			switch {
			case lrp.l1Stack[lrp.l1Iter].tokenType == Term:
				lrp.returnOnTerm()
				continue
			case lrp.l1Stack[lrp.l1Iter].tokenType == NTerm &&
				lrp.l1Stack[lrp.l1Iter].numOfAlternative < lrp.l1Stack[lrp.l1Iter].countAlternative:
				lrp.testAlternative()
				continue
			case lrp.l1Stack[lrp.l1Iter].tokenType == NTerm &&
				lrp.l1Stack[lrp.l1Iter].numOfAlternative >= lrp.l1Stack[lrp.l1Iter].countAlternative:
				if lrp.l1Stack[lrp.l1Iter].token == lrp.grammar.startSymbol && lrp.inIter == 0 {
					return fmt.Errorf("The input string does not belong to the grammar ")
				} else {
					lrp.returnNonTerm()
					continue
				}
			}

		case end:
			lrp.printOutput()
			return nil
		}
	}
}
