package automaton

import "automaton-builder/decode"

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

func NewAutomaton(ad *decode.AutomatonDescription) *Automaton {
	a := new(Automaton)

	// for _, v := range ad.States {
	// 	if v.Initial {
	// 		transitions = [ad.Transitions]

	// 		a.initialState = &State{
	// 			isFinal: v.Final,
	// 			transitions: ,
	// 		}
	// 	}
	// }

	for i := range ad.States {
		if ad.States[i].Initial {

			symbol := ad.Transitions[i].Symbol
			to := ad.Transitions[i].To

			ts = append(ts, &Transition{symbol: symbol, to: &State{}})
		}
	}
}

func getStateTransitions(stateName string, ad *decode.AutomatonDescription) []*Transition {
	ts := make([]*Transition, 5)

}
