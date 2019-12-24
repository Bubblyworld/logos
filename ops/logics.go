package ops

import "errors"

// Logic is an enum of supported formal logics.
type Logic int

const (
	logicStart         Logic = 0
	LogicPropositional Logic = 1
	logicEnd           Logic = 2
)

var (
	logicIDs = map[Logic]string{
		LogicPropositional: "propositional",
	}

	logicNames = map[Logic]string{
		LogicPropositional: "Propositional Logic",
	}

	logicDescs = map[Logic]string{
		LogicPropositional: `Propositional logic is the calculus and semantics of atomic propositions,
which have a truth value independent of the objects they refer to. It can also
be viewed as a representation of boolean algebra, as the permissible logical 
connectives are all truth-functional and hence isomorphic to boolean functions.`,
	}
)

func (l Logic) ID() string {
	id, ok := logicIDs[l]
	if !ok {
		panic("tried to find ID for invalid logic")
	}

	return id
}

func (l Logic) Name() string {
	name, ok := logicNames[l]
	if !ok {
		panic("tried to find name for invalid logic")
	}

	return name
}

func (l Logic) Description() string {
	desc, ok := logicDescs[l]
	if !ok {
		panic("tried to find description for invalid logic")
	}

	return desc
}

func ToLogic(id string) (Logic, error) {
	for l, i := range logicIDs {
		if i == id {
			return l, nil
		}
	}

	return 0, errors.New("no available logic with ID '" + id + "'")
}

func SupportedLogics() []Logic {
	var ll []Logic
	for i := logicStart + 1; i < logicEnd; i++ {
		ll = append(ll, i)
	}

	return ll
}
