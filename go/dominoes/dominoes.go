package dominoes

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return input, true
	}
	if len(input) == 1 {
		return input, input[0][0] == input[0][1]
	}

	head := input[0]
	for i, current := range input {
		if i == 0 {
			continue
		}

		if head[1] == current[1] {
			current[0], current[1] = current[1], current[0]
		}

		if head[1] == current[0] {
			// Move the pair to start of the new chain, compact it because we only need the head and tail.
			compacted := Domino([2]int{head[0], current[1]})
			remaining := []Domino{compacted}
			remaining = append(remaining, input[1:i]...)
			remaining = append(remaining, input[i+1:]...)

			chain, ok := MakeChain(remaining)
			if ok {
				fullChain := append([]Domino{head, current}, chain[1:]...)
				return fullChain, true
			}
		}
	}
	return nil, false
}
