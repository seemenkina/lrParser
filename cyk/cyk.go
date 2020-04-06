package cyk

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seemenkina/lrParser/grammar"
)

type MatrixIndex struct {
	X int
	Y int
}

type CYK struct {
	grammar grammar.Grammar
	input   string
	matrix  map[MatrixIndex]string
	output  string
}

func (c *CYK) NewCYK(gr grammar.Grammar, in string) {
	c.matrix = make(map[MatrixIndex]string)
	c.grammar = gr
	c.input = in
}

func (c *CYK) findTerm(symbol string) string {
	res := ""
	for _, rule := range c.grammar.Rules {
		if rule.RSymbol == symbol {
			res = fmt.Sprintf("%s%s", res, rule.LSymbol)
		}
	}
	return res
}

func (c *CYK) getMatrixVal(i, j int) string {
	return c.matrix[MatrixIndex{i, j}]
}

func makeOneSet(str string) []string {
	var res []string
	for i := 0; i < len(str); i++ {
		reStr := fmt.Sprintf("%s", string(str[i]))
		res = append(res, reStr)
	}
	return res
}

func makeSet(str1, str2 string) []string {
	if len(str1) == 0 {
		return makeOneSet(str2)
	} else if len(str2) == 0 {
		return makeOneSet(str1)
	}
	var res []string
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			reStr := fmt.Sprintf("%s%s", string(str1[i]), string(str2[j]))
			res = append(res, reStr)
		}
	}
	return res
}

func (c *CYK) StartCYK() bool {
	for i := 0; i < len(c.input); i++ {
		val := c.findTerm(string(c.input[i]))
		c.matrix[MatrixIndex{i, i}] = val
		fmt.Println(fmt.Sprintf("{%d %d}: %s", i, i, c.matrix[MatrixIndex{i, i}]))
	}

	for l := 1; l < len(c.input); l++ {
		for i := 0; i < len(c.input)-l; i++ {
			j := i + l
			var strSet []string
			res := ""

			for k := 0; k < j; k++ {
				val1 := c.getMatrixVal(i, i+k)
				val2 := c.getMatrixVal(i+k+1, j)
				set := makeSet(val1, val2)
				for _, s := range set {
					strSet = append(strSet, s)
				}
			}

			for _, ss := range strSet {
				val := c.findTerm(ss)
				for _, v := range val {
					if !strings.Contains(res, string(v)) {
						res = fmt.Sprintf("%s%s", res, string(v))
					}
				}
			}

			c.matrix[MatrixIndex{i, j}] = res
			fmt.Println(fmt.Sprintf("{%d %d}: %s", i, j, c.matrix[MatrixIndex{i, j}]))

		}

	}

	res := c.getMatrixVal(0, len(c.input)-1)
	if strings.Contains(res, c.grammar.StartSymbol) {
		return true
	}
	return false
}

func (c *CYK) findTermRule(sym string) int {
	for i, r := range c.grammar.Rules {
		if r.LSymbol == sym && len(r.RSymbol) == 1 {
			return i
		}
	}
	return -1
}

func (c *CYK) findNTermRules(sym string) map[int]string {
	ans := make(map[int]string)
	for i, r := range c.grammar.Rules {
		if r.LSymbol == sym {
			ans[i] = r.RSymbol
		}
	}
	return ans
}

func (c *CYK) PrintOut() {
	c.gen(0, len(c.input)-1, c.grammar.StartSymbol)
	fmt.Println(c.output)
}

func (c *CYK) gen(i, j int, A string) {
	if j == i {
		if num := c.findTermRule(A); num != -1 {
			c.output = fmt.Sprintf("%s %s", c.output, strconv.Itoa(num))
			return
		}
	}

	maps := c.findNTermRules(A)
	for k := j - 1; k >= 0; k-- {
		val, _ := c.matrix[MatrixIndex{i, i + k}]
		val2, _ := c.matrix[MatrixIndex{i + k + 1, j}]
		possible := makeSet(val, val2)
		for _, value := range possible {
			for key, val := range maps {
				if val == value {
					c.output = fmt.Sprintf("%s %s", c.output, strconv.Itoa(key))
					c.gen(i, i+k, string(value[0]))
					c.gen(i+k+1, j, string(value[1]))
					return
				}
			}
		}
	}
}
