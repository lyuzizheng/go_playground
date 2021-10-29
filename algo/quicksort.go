package algo

import (
	"sync"
)


func swap(a *[]float64, i int, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

const maxNum = 1000

func QuickSort(a *[]float64, left int, right int) {
	if right <= left {
		return
	}
	if (right - left) < 30 {
		InsertionSort(a, left, right)
		return
	}
	idx := left
	pivotIndex := findPivot(a, left, right)
	swap(a, pivotIndex, right)
	for i := left; i < right; i++ {
		if (*a)[i] < (*a)[right] {
			swap(a, i, idx)
			idx++
		}
	}
	swap(a, idx, right)
	QuickSort(a, left, idx)
	QuickSort(a, idx+1, right)
	return
}

func ConcurrentQuickSort(a *[]float64, left int, right int, ) {
	if (right - left) < maxNum {
		QuickSort(a, left, right)
		return
	}
	idx := left
	pivotIndex := findPivot(a, left, right)
	swap(a, pivotIndex, right)
	for i := left; i < right; i++ {
		if (*a)[i] < (*a)[right] {
			swap(a, i, idx)
			idx++
		}
	}
	swap(a, idx, right)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ConcurrentQuickSort(a, left, idx)
	}()
	ConcurrentQuickSort(a, idx+1, right)
	wg.Wait()
	return
}

func findPivot(f *[]float64, left int, right int)int{
	a := (*f)[left]
	b := (*f)[(left+right) >> 1]
	c := (*f)[right]
	if a < b {
		switch {
		case b < c:
			return (left+right) >> 1
		case a < c:
			return right
		default:
			return left
		}
	}
	switch {
	case a < c:
		return left
	case b < c:
		return right
	default:
		return (left+right) >> 1
	}
}


func InsertionSort(v *[]float64, left int, right int) {
	for i := left; i <= right; i++ {
		for j := i; j > 0 && (*v)[j-1] > (*v)[j]; j-- {
			swap(v, j, j-1)
		}
	}
}
