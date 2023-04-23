package automaton_test

// import (
// 	"automaton-builder/automaton"
// 	"testing"
// )

// func BuildAM(data [][]string) [][]*string {
// 	am := make([][]*string, 0)

// 	for i := range data {
// 		am = append(am, make([]*string, 0))

// 		for j := range data[i] {
// 			am[i] = append(am[i], &data[i][j])
// 		}
// 	}

// 	return am
// }

// func TestNewDFA(t *testing.T) {
// 	var am [][]*string
// 	var err error

// 	// It should create a new automaton
// 	am = BuildAM([][]string{
// 		{"a", "b"},
// 		{"a", "a"},
// 	})

// 	_, err = automaton.NewDFA(am, 0, []int{0})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// It should throw error if the matrix is not quadratic
// 	am = BuildAM([][]string{
// 		{"a", "b"},
// 		{"a", "a"},
// 	})

// 	_, err = automaton.NewDFA(am, 0, []int{0})
// 	if err == nil {
// 		t.Error("Expected: error", "Received: Automaton")
// 	}

// 	// am := automaton.NewAdjacencyMatrix("- 0 1\n- 1 -\n- 0 -\n")
// }
