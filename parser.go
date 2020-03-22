package main

type L1Token struct {
	token     string
	tokenType int
	// for non terminal symbol
	countAlternative int // number of alternative in rules
	firstN           int // first alternative in rules
}

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

type State struct {
	st int
}

type LRParser struct {
	grammar Grammar
	input   string
	l1Stack []L1Token
	l2Stack []L2Token
	state   State
}

func (lrp *LRParser) newLRParser(gr Grammar, in string) {
	lrp.grammar = gr
	lrp.input = in
	lrp.state.st = normal
}
