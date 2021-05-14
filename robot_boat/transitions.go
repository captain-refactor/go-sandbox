package robot_boat

type Transitions struct {
	transitions [][2]State
}

func isValidTransition(a, b State) bool {
	if a.Robot == b.Robot {
		return false
	}
	if a.Robot == Far {
		// assure that we will deal with near to far transitions
		return isValidTransition(b, a)
	}
	if !validate(a) || !validate(b) {
		return false
	}
	deltaCount := 0
	for _, pair := range [3][2]Position{
		{a.Grain, b.Grain},
		{a.Hen, b.Hen},
		{a.Fox, b.Fox},
	} {
		if pair[0] != pair[1] {
			deltaCount++
			if pair[1] == Near {
				return false
			}
		}
	}

	// there are only two seats, so we cannot take three things
	if deltaCount > 1 {
		return false
	}
	return true
}

func (t Transitions) Transitions() [][2]State {
	return t.transitions
}

func NewTransitions() Transitions {
	return Transitions{}
}

// AddTransition expects valid transition
func (t Transitions) AddTransition(a, b State) Transitions {
	if a.Robot == Far {
		return t.AddTransition(b, a)
	}
	newTransition := [2]State{a, b}
	for _, transition := range t.transitions {
		if transition == newTransition {
			return t
		}
	}
	t.transitions = append(t.transitions, newTransition)
	return t
}

func (t Transitions) GetTransitions(s State) (states []State) {
	keyIndex := 0
	targetIndex := 1
	if s.Robot == Far {
		keyIndex = 1
		targetIndex = 0
	}
	for _, transition := range t.transitions {
		if transition[keyIndex] == s {
			states = append(states, transition[targetIndex])
		}
	}
	return
}

func CreateValidTransitions(states []State) Transitions {
	validTransitions := NewTransitions()
	for _, a := range states {
		for _, b := range states {
			if isValidTransition(a, b) {
				validTransitions = validTransitions.AddTransition(a, b)
			}
		}
	}
	return validTransitions
}
