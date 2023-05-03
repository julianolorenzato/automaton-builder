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

// func removeDuplicates[T string | int](s []T) bool {
// 	duplicates := make(map[T]int)

// 	for i := range s {

// 	}

// }
