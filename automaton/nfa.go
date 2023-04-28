package automaton

type NFA struct {
	Name         string
	Alphabet     []string
	CurrentState *State
	InitialState *State
	States       []State
	Transitions  map[Pair][]*State
}
