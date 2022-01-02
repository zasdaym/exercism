// Package secret provides Secret Handshake solution.
package secret

// events contains possible events in a secret handshake.
var events = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake creates sequence of events for a secret handshake given a code number.
func Handshake(code uint) []string {
	var result []string
	for i, event := range events {
		mask := uint(1 << i)
		if code&mask == mask {
			result = append(result, event)
		}
	}
	if code&16 == 16 {
		reverse(result)
	}
	return result
}

// reverse reverses a slice of strings.
func reverse(nums []string) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
