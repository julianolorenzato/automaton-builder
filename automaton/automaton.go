package automaton

type Automaton interface {
	Transition()
	Perform()
}

type State struct {
	name    string
	isFinal bool
}

type Pair struct {
	symbol *string
	from   *State
}
