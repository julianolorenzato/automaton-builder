package main

import (
	"fmt"

	"automaton-builder/automaton"
)

func main() {
	m, err := automaton.NewAdjacencyMatrix()

	if err != nil {
		fmt.Println(err)
	}

	for i := range m {
		for _, v := range m[i] {
			if v != nil {
				fmt.Println(*v)
			} else {
				fmt.Println("NULL")
			}
		}
	}
}
