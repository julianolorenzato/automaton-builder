package decode

import (
	"encoding/json"
	"io/ioutil"
)

type AutomatonDescription struct {
	Name        string
	Alphabet    []string
	States      []StateDescription
	Transitions []TransitionDescription
}

type StateDescription struct {
	Name    string
	Initial bool
	Final   bool
}

type TransitionDescription struct {
	From   string
	To     string
	Symbol string
}

// Revisar esta função!
// Na verdade, revisar todo esse package,
// talvez não seja performático decodificar json para uma "description" antes de passar para a estrutura do automato
func Decode() (*AutomatonDescription, error) {
	var ad AutomatonDescription

	raw, err := ioutil.ReadFile("example.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &ad)
	if err != nil {
		return nil, err
	}

	err = validateChain(&ad)
	if err != nil {
		return nil, err
	}

	return &ad, nil
}
