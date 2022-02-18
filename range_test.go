package RangeList

import (
	"testing"
)

func TestRangeList(t *testing.T) {
	rl := RangeList{rl: []section{}}
	rl.Add([2]int{1, 5})
	rl.Print()
	// Should display: [1, 5)
	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{2, 4})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 11})
	rl.Print()
	//Should display: [1, 8) [11, 21)
	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)
	rl.Remove([2]int{3, 19})
	rl.Print() // Should display: [1, 3) [19, 21)
}