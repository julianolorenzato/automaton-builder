package main

import (
	"automaton-builder/automaton"
	"fmt"
	"os"
)

func main() {
	am := automaton.NewAdjacencyMatrix()

	a, err := automaton.NewAutomaton(*am, 2, []int{0, 1, 2})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(a)
}
