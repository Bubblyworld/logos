package internal

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/bubblyworld/logos/ops"
)

var lex = lexer.Must(ebnf.New(`
    Conjunction = "` + ops.PROP_CONJ + `" .
    Disjunction = "` + ops.PROP_DISJ + `" .
    Implication = "` + ops.PROP_IMPL + `" .
    Not         = "` + ops.PROP_NOT + `" .
    LBracket    = "(" .
    RBracket    = ")" .
    Whitespace  = " " | "\t" | "\n" | "\r" .
    Atom        = { alpha } .
    alpha       = "a"…"z" | "A"…"Z" .
`))

var parser = participle.MustBuild(
	new(Formula),
	participle.Lexer(lex),
	participle.Elide("Whitespace"),
)

type Formula struct {
	Atomic     *AtomicFormula     `  @@`
	Negation   *NegationFormula   `| @@`
	Connective *ConnectiveFormula `| @@`
}

type AtomicFormula struct {
	Atom *string `@Atom`
}

type NegationFormula struct {
	Formula *Formula `Not @@`
}

type ConnectiveFormula struct {
	LFormula   *Formula    `LBracket @@`
	Connective *Connective `@@`
	RFormula   *Formula    `@@ RBracket`
}

type Connective struct {
	Conjunction *string `  @Conjunction`
	Disjunction *string `| @Disjunction`
	Implication *string `| @Implication`
}

func Parse(formula string) (*Formula, error) {
	var res Formula
	if err := parser.ParseString(formula, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
