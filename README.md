# golang-rangeList

- Task: Implement a struct named 'RangeList'
- A pair of integers define a range, for example: [1, 5). This range includes integers: 1, 2, 3, and 4.
- A range list is an aggregate of these ranges: [1, 5), [10, 11), [100, 201)
- NOTE: Feel free to add any extra member variables/functions you like.

```
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

```


