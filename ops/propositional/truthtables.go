package propositional

import (
	"sort"
	"strings"
)

// TruthTable represents the result of evaluating assignments against
// a formula.
type TruthTable map[Model]Value

// GetTruthTable returns the formula's truth table for all possible assigments
// of it's variables to true/false according to the given index.
func GetTruthTable(f Formula, index []string) TruthTable {
	tt := make(TruthTable)
	if len(index) == 0 {
		return tt
	}

	for _, a := range allModels(index) {
		tt[a] = Evaluate(f, a, index)
	}

	return tt
}

func (tt TruthTable) String() string {
	var al []Model
	for a := range tt {
		al = append(al, a)
	}

	sort.Slice(al, func(i, j int) bool {
		return al[i] < al[j]
	})

	var sl []string
	for _, a := range al {
		sl = append(sl, string(a)+":"+string(tt[a]))
	}

	return "[" + strings.Join(sl, " ") + "]"
}

func allModels(index []string) []Model {
	if len(index) <= 0 {
		return nil
	}

	if len(index) == 1 {
		return []Model{"0", "1"}
	}

	var all []Model
	for _, a := range allModels(index[1:]) {
		all = append(all, "0"+a)
		all = append(all, "1"+a)
	}

	return all
}

// Evaluate returns the result of assigning the identifiers in a formula
// the given truth values. Models should be indexed according to the
// given list of identifiers.
func Evaluate(f Formula, assign Model, index []string) Value {
	var fn func(a, b Value) Value
	switch f.Type {
	case Atomic:
		i := indexOf(f.Identifier, index)
		return Value(assign[i])

	case Negation:
		return evalNot(Evaluate(f.Subformulae[0], assign, index))

	case Conjunction:
		fn = evalConjunction

	case Disjunction:
		fn = evalDisjunction

	case Implication:
		fn = evalImplication

	default:
		// Sanity check - should never happen.
		panic("tried to evaluate malformed formula")
	}

	l := Evaluate(f.Subformulae[0], assign, index)
	r := Evaluate(f.Subformulae[1], assign, index)
	return fn(l, r)
}

func evalNot(a Value) Value {
	return ToValue(!ToBool(a))
}

func evalConjunction(a, b Value) Value {
	return ToValue(ToBool(a) && ToBool(b))
}

func evalDisjunction(a, b Value) Value {
	return ToValue(ToBool(a) || ToBool(b))
}

func evalImplication(a, b Value) Value {
	return ToValue((ToBool(a) == false) || (ToBool(b) == true))
}

func indexOf(id string, idl []string) int {
	for i, idlID := range idl {
		if id == idlID {
			return i
		}
	}

	return -1
}
