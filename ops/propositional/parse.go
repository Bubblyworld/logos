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
		return not(convert(f.Negation.Formula))
	}

	if f.Connective != nil {
		fn := getConstructor(f.Connective.Connective)

		return fn(
			convert(f.Connective.LFormula),
			convert(f.Connective.RFormula),
		)
	}

	// Sanity check - this should never happen.
	panic("malformed formula returned by propositional parse")
}

func getConstructor(c *internal.Connective) func(Formula, Formula) Formula {
	if c.Conjunction != nil {
		return and
	}

	if c.Disjunction != nil {
		return or
	}

	if c.Implication != nil {
		return implies
	}

	// Sanity check - should never happen.
	panic("malformed connective returned by propositional parser")
}
