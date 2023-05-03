package decode

import (
	"encoding/csv"
	"fmt"
	"os"
)

func DecodeCSV() [][]string {
	m, err := csv.NewReader(os.Stdin).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return m
}

// func validateQuadratic