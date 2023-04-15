package automaton

// import "automaton-builder/decode"

type Automaton struct {
	initialState *State
}

type Transition struct {
	symbol string
	to     *State
}

type State struct {
	isFinal     bool
	transitions []*Transition
}

// func NewAutomaton(ad *decode.AutomatonDescription) *Automaton {
// 	a := new(Automaton)

// 	for i := range am {
// 		for j := range am[i] {

// 		}
// 	}
// }
