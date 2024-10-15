package yacht

func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return count(dice, 1)
	case "twos":
		return count(dice, 2) * 2
	case "threes":
		return count(dice, 3) * 3
	case "fours":
		return count(dice, 4) * 4
	case "fives":
		return count(dice, 5) * 5
	case "sixes":
		return count(dice, 6) * 6
	case "full house":
		return fullHouse(dice)
	case "four of a kind":
		return fourOfAKind(dice)
	case "little straight":
		return littleStraight(dice)
	case "big straight":
		return bigStraight(dice)
	case "choice":
		return choice(dice)
	case "yacht":
		return yacht(dice)
	default:
		// Unknown category
		return -1
	}
}

func count(nums []int, target int) int {
	var result int
	for _, num := range nums {
		if num == target {
			result++
		}
	}
	return result
}

func fullHouse(nums []int) int {
	var sum int
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num] += 1
		sum += num
	}

	var hasDouble, hasTriple bool
	for _, freq := range frequencies {
		switch freq {
		case 2:
			hasDouble = true
		case 3:
			hasTriple = true
		}
	}

	if hasDouble && hasTriple {
		return sum
	} else {
		return 0
	}
}

func fourOfAKind(nums []int) int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num] += 1
		if frequencies[num] == 4 {
			return num * 4
		}
	}
	return 0
}

func littleStraight(nums []int) int {
	checks := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
	}
	for _, num := range nums {
		checks[num] = true
	}
	for _, check := range checks {
		if !check {
			return 0
		}
	}
	return 30
}

func bigStraight(nums []int) int {
	checks := map[int]bool{
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
	}
	for _, num := range nums {
		checks[num] = true
	}
	for _, check := range checks {
		if !check {
			return 0
		}
	}
	return 30
}

func choice(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

func yacht(nums []int) int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num] += 1
		if frequencies[num] == 5 {
			return 50
		}
	}
	return 0
}
