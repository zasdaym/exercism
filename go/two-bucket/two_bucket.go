package twobucket

import "errors"

var (
	ErrInvalidBucketName             = errors.New("invalid bucket name")
	ErrInvalidBucketSize             = errors.New("invalid bucket size")
	ErrInvalidGoalAmount             = errors.New("invalid goal amount")
	ErrGoalAmountMoreThanBucketSizes = errors.New("goal amount exceeds bucket sizes")
	ErrNoSolution                    = errors.New("no solution")
)

type Bucket struct {
	size, amount int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne == 0 || sizeBucketTwo == 0 {
		return "", 0, 0, ErrInvalidBucketSize
	}
	if goalAmount == 0 {
		return "", 0, 0, ErrInvalidGoalAmount
	}
	if goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo {
		return "", 0, 0, ErrGoalAmountMoreThanBucketSizes
	}
	if goalAmount%gcd(sizeBucketOne, sizeBucketTwo) != 0 {
		return "", 0, 0, ErrNoSolution
	}

	var start, other int
	switch startBucket {
	case "one":
		start, other = 0, 1
	case "two":
		start, other = 1, 0
	default:
		return "", 0, 0, ErrInvalidBucketName
	}

	buckets := [2]Bucket{
		{size: sizeBucketOne},
		{size: sizeBucketTwo},
	}

	var steps int
	for buckets[0].amount != goalAmount && buckets[1].amount != goalAmount {
		steps++
		switch {
		case buckets[start].amount == 0:
			buckets[start].amount = buckets[start].size
		case buckets[other].size == goalAmount:
			buckets[other].amount = goalAmount
		case buckets[other].amount == buckets[other].size:
			buckets[other].amount = 0
		default:
			pour := min(buckets[other].size-buckets[other].amount, buckets[start].amount)
			buckets[start].amount -= pour
			buckets[other].amount += pour
		}
	}

	if buckets[0].amount == goalAmount {
		return "one", steps, buckets[1].amount, nil
	} else {
		return "two", steps, buckets[0].amount, nil
	}
}
