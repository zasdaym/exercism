package dominoes

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return nil, true
	}
	if len(input) == 1 {
		return input, input[0][0] == input[0][1]
	}

	unused := make([]Domino, len(input))
	copy(unused, input)

	for i := 0; i < len(unused); i++ {
		first := unused[i]
		chain := []Domino{first}
		
		unused[i] = unused[len(unused)-1]
		remaining := unused[:len(unused)-1]
		
		if result, ok := backtrack(chain, remaining, first[0], first[1]); ok {
			return result, true
		}
		if result, ok := backtrack(chain, remaining, first[1], first[0]); ok {
			chain[0] = Domino{first[1], first[0]}
			return result, true
		}
		
		unused[i] = first
	}
	
	return nil, false
}

func backtrack(chain []Domino, unused []Domino, start, end int) ([]Domino, bool) {
	if len(unused) == 0 {
		return chain, start == end
	}
	
	for i := 0; i < len(unused); i++ {
		dom := unused[i]
		
		if end == dom[0] {
			newChain := append([]Domino{}, chain...)
			newChain = append(newChain, dom)
			
			unused[i] = unused[len(unused)-1]
			newUnused := unused[:len(unused)-1]
			
			if result, ok := backtrack(newChain, newUnused, start, dom[1]); ok {
				return result, true
			}
			
			unused[i] = dom
		}
		
		if end == dom[1] {
			newChain := append([]Domino{}, chain...)
			newChain = append(newChain, Domino{dom[1], dom[0]})
			
			unused[i] = unused[len(unused)-1]
			newUnused := unused[:len(unused)-1]
			
			if result, ok := backtrack(newChain, newUnused, start, dom[0]); ok {
				return result, true
			}
			
			unused[i] = dom
		}
	}
	
	return nil, false
}
