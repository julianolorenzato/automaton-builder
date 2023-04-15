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
		validateFinalStates(ad),
		validateStatesName(ad),
	)

	return errs
}

// Verify if the automaton description have exactly one initial state
func validateInitialStates(ad *AutomatonDescription) error {
	quantity := 0

	for _, v := range ad.States {
		if v.Initial {
			quantity++
		}
	}

	if quantity == 0 {
		return errors.New("automaton description doesnt have an initial state")
	} else if quantity > 1 {
		return errors.New("automaton description have more than one initial state")
	}

	return nil
}

// Verify if the automaton description have at least one final state
func validateFinalStates(ad *AutomatonDescription) error {
	quantity := 0

	for _, v := range ad.States {
		if v.Final {
			quantity++
		}
	}

	if quantity == 0 {
		return errors.New("automaton description doesnt have at least final states")
	}

	return nil
}

// Verify if the automaton does not have a duplicate state name
func validateStatesName(ad *AutomatonDescription) error {
	eachQuantity := map[string]int{}

	for _, v := range ad.States {
		eachQuantity[v.Name]++
	}

	for k, v := range eachQuantity {
		if v > 1 {
			return errors.New(`automaton description have more than one state name "` + k + `"`)
		}
	}

	return nil
}
