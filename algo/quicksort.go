package algo

import (
	"sync"
)

func swap(a *[]float64, i int, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

const (
	concurrentLimit = 2000
	quicksortLimit  = 30
)

func QuickSort(f *[]float64, left int, right int) {
	if (right - left) < quicksortLimit {
		InsertionSort(f, left, right)
		return
	}
	idx := left
	//pivotIndex := findPivot(f, left, right)
	pivotIndex := left
	swap(f, pivotIndex, right)
	for i := left; i < right; i++ {
		if (*f)[i] < (*f)[right] {
			swap(f, i, idx)
			idx++
		}
	}
	swap(f, idx, right)
	QuickSort(f, left, idx)
	QuickSort(f, idx+1, right)
	return
}

func ConcurrentQuickSort(f *[]float64, left int, right int) {
	if (right - left) < concurrentLimit {
		QuickSort(f, left, right)
		return
	}
	idx := left
	//pivotIndex := findPivot(f, left, right)
	pivotIndex := left
	swap(f, pivotIndex, right)
	for i := left; i < right; i++ {
		if (*f)[i] < (*f)[right] {
			swap(f, i, idx)
			idx++
		}
	}
	swap(f, idx, right)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ConcurrentQuickSort(f, left, idx)
		wg.Done()
	}()
	ConcurrentQuickSort(f, idx+1, right)
	wg.Wait()
	return
}

func findPivot(f *[]float64, left int, right int) int {
	middle := (left + right) >> 1
	a := (*f)[left]
	b := (*f)[middle]
	c := (*f)[right]
	if a < b {
		switch {
		case b < c:
			return middle
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
		return middle
	}
}

func InsertionSort(f *[]float64, left int, right int) {
	for i := left; i <= right; i++ {
		for j := i; j > 0 && (*f)[j-1] > (*f)[j]; j-- {
			swap(f, j, j-1)
		}
	}
}
