package main

import (
	"automaton-builder/automaton"
	"encoding/csv"
	"fmt"
	"os"
	// "encoding/json"
)

func main() {
	// am := automaton.NewAdjacencyMatrix("- 0 1\n- 1 -\n- 0 -\n")

	// var a automaton.Automaton
	// var err error

	// a, err = automaton.NewDFA(*am, 2, []int{0, 1, 2})
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// }

	// // name, states, finalStates := a.GetInfo()

	// // fmt.Println(name)
	// // fmt.Println(states)
	// // fmt.Println(finalStates)

	// s, ok, err := a.Perform([]string{"1"})

	// fmt.Println(s, ok, err)

	file, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	m, _ := csv.NewReader(file).ReadAll()

	a, err := automaton.NewDFA(m, []int{3})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*a)

	isValid, _ := a.Perform([]string{"0", "1", "1", "0", "1"})

	fmt.Println(*a.CurrentState)
	fmt.Println(isValid)
}
