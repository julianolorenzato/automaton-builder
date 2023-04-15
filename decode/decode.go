package decode

import (
	"encoding/json"
	"errors"
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

// Perform all validations and return any errors found
func validateChain(ad *AutomatonDescription) error {
	errs := errors.Join(
		validateInitialStates(ad),
	)

	return errs
}

// Verify if the automaton have exactly one initial state
func validateInitialStates(ad *AutomatonDescription) error {
	quantity := 0

	for _, v := range ad.States {
		if v.Initial {
			quantity++
		}
	}

	if quantity == 0 {
		return errors.New("input doesnt have an initial state")
	} else if quantity > 1 {
		return errors.New("input have more than one initial state")
	}

	return nil
}
