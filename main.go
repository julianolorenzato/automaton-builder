package main

import (
	"automaton-builder/automaton"
	"automaton-builder/decode"
	"fmt"
)

func main() {
	m := decode.DecodeCSV()

	a, err := automaton.NewDFA(m)
	if err != nil {
		fmt.Println(err)
	}

	isValid, err := a.Perform("0", "a", "1", "0", "1", "1", "1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(isValid)
}
