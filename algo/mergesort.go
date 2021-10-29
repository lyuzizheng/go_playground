package algo

import (
	"sync"
)



const max = 50

func mergesortParallel(s []float64) []float64{
	len := len(s)

	if len > 1 {
		if len <= max { // Sequential
			return mergesort(s)
		} else { // Parallel
			middle := len / 2
			var left []float64
			var right []float64

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				right = mergesortParallel(s[:middle])
			}()

			left = mergesortParallel(s[middle:])

			wg.Wait()
			return merge(left, right)
		}
	}
	return s

}

func mergesort(s []float64) []float64 {
	if len(s) > 1 {
		middle := len(s) / 2
		return merge(mergesort(s[:middle]), mergesort(s[middle:]))
	}
	return s
}

func merge(left, right []float64) (slice []float64) {
	size, i, j := len(left)+len(right), 0, 0
	slice = make([]float64, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}