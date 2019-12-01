package state

// State contains all the persistent, mutable state required by logos for proof
// generation and checking.
type State struct {
}

func New() (*State, error) {
	return &State{}, nil
}
