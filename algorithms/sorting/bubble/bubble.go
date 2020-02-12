package bubble

import "fmt"

// Sort func
func Sort(in []int) []int {

	max := len(in)

	for i := 0; i < max; i++ {
		fmt.Println("i ", i)

		var next, current, nIdx, cIdx int
		for j := 0; j < max; j++ {
			cIdx = j
			nIdx = j + 1

			if cIdx < nIdx {
				current = in[cIdx]
				next = in[nIdx]
				fmt.Printf("current %d \t  next %d  \n", current, next)

				if current > next {
					fmt.Printf("swapping current %d \t with next %d  \n", current, next)
					tmp := next
					in[nIdx] = current
					in[cIdx] = tmp
					fmt.Printf("swapped current %d \t with next %d  \n", in[cIdx], in[nIdx])
				}
			}

		}

		// we have the highest value in the last slot, no need to test again
		max--
	}
	return in
}
