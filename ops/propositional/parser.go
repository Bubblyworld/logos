package propositional

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/bubblyworld/logos/ops"
)

// TODO(guy): Handle precedence and bracket ellision.
// TODO(guy): This stuff should go in an internal package and we should return
//            a nicer data structure from Parse().

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

func (f *Formula) String() string {
	if f.Atomic != nil {
		return f.Atomic.String()
	}

	if f.Negation != nil {
		return f.Negation.String()
	}

	if f.Connective != nil {
		return f.Connective.String()
	}

	return "malformed formula"
}

type AtomicFormula struct {
	Atom *string `@Atom`
}

func (af *AtomicFormula) String() string {
	if af.Atom == nil {
		return "malformed atomic formula"
	}

	return *af.Atom
}

type NegationFormula struct {
	Formula *Formula `Not @@`
}

func (nf *NegationFormula) String() string {
	if nf.Formula == nil {
		return "malformed negation formula"
	}

	return ops.PROP_NOT + nf.Formula.String()
}

type ConnectiveFormula struct {
	LFormula   *Formula    `LBracket @@`
	Connective *Connective `@@`
	RFormula   *Formula    `@@ RBracket`
}

func (cf *ConnectiveFormula) String() string {
	if cf.LFormula == nil || cf.Connective == nil || cf.RFormula == nil {
		return "malformed connective formula"
	}

	return "(" + cf.LFormula.String() + " " + cf.Connective.String() +
		" " + cf.RFormula.String() + ")"
}

type Connective struct {
	Conjunction *string `  @Conjunction`
	Disjunction *string `| @Disjunction`
	Implication *string `| @Implication`
}

func (c *Connective) String() string {
	if c.Conjunction != nil {
		return *c.Conjunction
	}

	if c.Disjunction != nil {
		return *c.Disjunction
	}

	if c.Implication != nil {
		return *c.Implication
	}

	return "malformed connective"
}

func Parse(formula string) (*Formula, error) {
	var res Formula
	if err := parser.ParseString(formula, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
