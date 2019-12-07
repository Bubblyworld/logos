package propositional

// GetConjunctiveNormalForm returns a logically equivalent formula that
// consists entirely of a conjunction of disjunctive clauses.
func GetConjunctiveNormalForm(f Formula) Formula {
	return distribute(simplify(f, false))
}

// simplify returns a logically equivalent formula that contains no
// implications and which only has negations of atomic formulae.
func simplify(f Formula, negating bool) Formula {
	switch f.Type {
	case Atomic:
		if negating {
			return not(f)
		}

		return f

	case Negation:
		return simplify(f.fst(), !negating)

	// A -> B <==> B | !A
	case Implication:
		return or(
			simplify(f.snd(), negating),
			simplify(f.fst(), !negating),
		)

	// !(A ^ B) <==> (!A) | (!B)
	case Disjunction:
		if negating {
			return and(
				simplify(f.fst(), true),
				simplify(f.snd(), true),
			)
		}

		return or(
			simplify(f.fst(), false),
			simplify(f.snd(), false),
		)

	// !(A | B) <==> (!A) ^ (!B)
	case Conjunction:
		if negating {
			return or(
				simplify(f.fst(), true),
				simplify(f.snd(), true),
			)
		}

		return and(
			simplify(f.fst(), false),
			simplify(f.snd(), false),
		)

	default:
		panic("tried to simplify malformed formula")
	}
}

// distribute returns a logically equivalent formula in which disjunctions
// have been distributed over conjunctions. Note that the input formula MUST
// already be simplified.
func distribute(f Formula) Formula {
	switch f.Type {
	case Atomic:
		return f

	case Negation:
		return f

	case Conjunction:
		return and(
			distribute(f.fst()),
			distribute(f.snd()),
		)

	case Disjunction:
		fst := distribute(f.fst())
		snd := distribute(f.snd())

		if fst.Type != Conjunction && snd.Type != Conjunction {
			return or(fst, snd)
		}

		if fst.Type == Conjunction {
			fst, snd = snd, fst // snd is conjunction for consistency
		}

		// A | (B ^ C) <==> (A | B) ^ (A | C)
		return distribute(and(
			or(fst, snd.fst()),
			or(fst, snd.snd()),
		))

	default:
		panic("tried to distribute over non-simplified formula")
	}
}
