package propositional

import (
	"sort"

	"github.com/bubblyworld/logos/ops"
)

// TODO(guy): It's bad design to be extremely aggressive about panic()
// on malformed input while not providing a mechanism to produce correctly
// formed input (other than parsing, but that's bad programatically).

type FormulaType int

const (
	Atomic      FormulaType = 1
	Negation    FormulaType = 2
	Conjunction FormulaType = 3
	Disjunction FormulaType = 4
	Implication FormulaType = 5
)

// Formula represents a propositional formula as a k-ary logical connective
// along with k subformulas that represent the arguments to the connective.
// This approach is intended to make it easy to recurse on formulae.
type Formula struct {
	Type        FormulaType
	Subformulae []Formula

	// Identifier is only non-zero if the formula is atomic.
	Identifier string
}

var arities = map[FormulaType]int{
	Atomic:      0,
	Negation:    1,
	Conjunction: 2,
	Disjunction: 2,
	Implication: 2,
}

func (f Formula) Arity() int {
	return arities[f.Type]
}

// ListIdentifiers returns the list of unique atomic formula identifiers
// contained in the given formula. The result is guaranteed to be sorted
// lexicographically.
func (f Formula) ListIdentifiers() []string {
	if f.Arity() == 0 {
		return []string{f.Identifier}
	}

	idm := make(map[string]bool)
	for i := 0; i < f.Arity(); i++ {
		for _, id := range f.Subformulae[i].ListIdentifiers() {
			idm[id] = true
		}
	}

	var idl []string
	for id := range idm {
		idl = append(idl, id)
	}

	sort.Strings(idl)
	return idl
}

func (f Formula) String() string {
	var c string
	switch f.Type {
	case Atomic:
		return f.Identifier

	case Negation:
		return ops.PROP_NOT + f.Subformulae[0].String()

	case Conjunction:
		c = ops.PROP_CONJ

	case Disjunction:
		c = ops.PROP_DISJ

	case Implication:
		c = ops.PROP_IMPL

	default:
		return "unknown formula type"
	}

	return "(" + f.Subformulae[0].String() + " " + c +
		" " + f.Subformulae[1].String() + ")"
}
