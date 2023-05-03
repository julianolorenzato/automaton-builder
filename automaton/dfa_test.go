package automaton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDFA(t *testing.T) {
	// It should create a new DFA
	var m [][]string = [][]string{
		{"", "0", "1"},
		{"*S0", "S0", "S1"},
		{"S1", "S2", "S0"},
		{"S2", "S1", "S2"},
	}

	dfa, err := NewDFA(m)
	assert.Nil(t, err)
	assert.NotNil(t, dfa)
	assert.IsType(t, new(DFA), dfa)
	assert.Equal(t, []State{{"S0", true}, {"S1", false}, {"S2", false}}, dfa.States)
	assert.Equal(t, []string{"0", "1"}, dfa.Alphabet)

	// -------------------------------
	// It should not have an invalid state in the transitions
	m = [][]string{
		{"", "0", "1"},
		{"*S0", "S3", "S1"},
		{"S1", "S2", "S0"},
		{"S2", "S1", "S2"},
	}

	dfa, err = NewDFA(m)
	assert.Nil(t, dfa)
	assert.ErrorContains(t, err, "no state of name ")
}
