package decode

import "errors"

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
