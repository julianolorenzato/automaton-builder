package automaton

import (
	"automaton-builder/util"
	"errors"
	"fmt"
)

type DFA struct {
	Name         string
	Alphabet     []string
	CurrentState *State
	InitialState *State
	States       []State
	Transitions  map[Pair]*State
}

// Params: m - matrix
func NewDFA(m [][]string) (*DFA, error) {
	dfa := DFA{
		Transitions: make(map[Pair]*State),
	}

	// Add Alphabet Symbols to dfa (ignore first column)
	dfa.Alphabet = m[0][1:]

	// Remove first line after
	m = m[1:]

	// Add States to dfa (ignore first line)
	for i := range m {
		stateName := m[i][0]
		isFinal := stateName[0] == '*'

		if isFinal {
			stateName = stateName[1:]
		}

		dfa.States = append(dfa.States, State{
			name:    stateName,
			isFinal: isFinal,
		})
		// Remove first column after
		m[i] = m[i][1:]
	}

	dfa.InitialState = &dfa.States[0]
	dfa.CurrentState = &dfa.States[0]

	// Add Transitions to dfa
	for i := range m {
		for j := range m[i] {
			index := util.IndexOf(dfa.States, func(s State) bool {
				return m[i][j] == s.name
			})

			if index == -1 {
				return nil, errors.New("no state of name " + m[i][j])
			}

			dfa.Transitions[Pair{
				from:   &dfa.States[i],
				symbol: &dfa.Alphabet[j],
			}] = &dfa.States[index]
		}
	}

	return &dfa, nil
}

func (dfa *DFA) Transition(symbol string) error {
	var s *string

	// Get the adress of symbol in alphabet
	for i := range dfa.Alphabet {
		if symbol == dfa.Alphabet[i] {
			s = &dfa.Alphabet[i]
		}
	}
	if s == nil {
		return errors.New(`symbol "` + symbol + `" doest not belongs to the automaton alphabet`)
	}

	// Get the adress of the destination state
	to, ok := dfa.Transitions[Pair{
		symbol: s,
		from:   dfa.CurrentState,
	}]
	if !ok {
		return errors.New("no transitions from state " + dfa.CurrentState.name + " with symbol " + symbol)
	}

	dfa.CurrentState = to
	return nil
}

func (dfa *DFA) Perform(w ...string) (bool, error) {
	for _, v := range w {
		err := dfa.Transition(v)

		if err != nil {
			return false, err
		}
	}

	if dfa.CurrentState.isFinal {
		return true, nil
	} else {
		return false, errors.New("the word was read but the end state is not final")
	}
}

func (dfa *DFA) Reset() {
	dfa.CurrentState = dfa.InitialState
}

func (dfa *DFA) PrintInfo() {

	fmt.Println("==========================")
	fmt.Println("Initial State:")
	fmt.Println("--------------------------")
	fmt.Println(*dfa.InitialState)

	fmt.Println("==========================")
	fmt.Println("Current State:")
	fmt.Println("--------------------------")
	fmt.Println(*dfa.CurrentState)

	fmt.Println("==========================")
	fmt.Println("Alphabet:")
	fmt.Println("--------------------------")
	for i := range dfa.Alphabet {
		fmt.Println(`"` + dfa.Alphabet[i] + `"`)
	}

	fmt.Println("==========================")
	fmt.Println("States:")
	fmt.Println("--------------------------")
	for i := range dfa.States {
		fmt.Println(dfa.States[i].name)
	}

	fmt.Println("==========================")
	fmt.Println("Transitions:")
	fmt.Println("--------------------------")
	for k, v := range dfa.Transitions {
		fmt.Println(k.from.name, `"`+*k.symbol+`"`, v.name)
	}
	fmt.Println("==========================")
}
