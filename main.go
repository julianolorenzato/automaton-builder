package main

import (
	"automaton-builder/automaton"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	var err error

	m, err := csv.NewReader(os.Stdin).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

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
