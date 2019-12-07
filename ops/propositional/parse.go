package propositional

import (
	"github.com/bubblyworld/logos/ops/propositional/internal"
)

func Parse(formula string) (Formula, error) {
	f, err := internal.Parse(formula)
	if err != nil {
		return Formula{}, err
	}

	return convert(f), nil
}

func convert(f *internal.Formula) Formula {
	if f.Atomic != nil {
		return Formula{
			Type:       Atomic,
			Identifier: *f.Atomic.Atom,
		}
	}

	if f.Negation != nil {
		return Formula{
			Type: Negation,
			Subformulae: []Formula{
				convert(f.Negation.Formula),
			},
		}
	}

	if f.Connective != nil {
		return Formula{
			Type: getType(f.Connective.Connective),
			Subformulae: []Formula{
				convert(f.Connective.LFormula),
				convert(f.Connective.RFormula),
			},
		}
	}

	// Sanity check - this should never happen.
	panic("malformed formula returned by propositional parse")
}

func getType(c *internal.Connective) FormulaType {
	if c.Conjunction != nil {
		return Conjunction
	}

	if c.Disjunction != nil {
		return Disjunction
	}

	if c.Implication != nil {
		return Implication
	}

	// Sanity check - should never happen.
	panic("malformed connective returned by propositional parser")
}
