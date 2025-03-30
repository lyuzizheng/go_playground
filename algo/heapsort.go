package algo

func heapsort(items []int) {
	n := len(items)

	buildHeap(items, n)

	for i := n - 1; i >= 0; i-- {
		items[0], items[i] = items[i], items[0]
		heapify(items, i, 0)
	}

}

func buildHeap(items []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(items, n, i)
	}
}

func heapify(items []int, n int, index int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < n && items[left] > items[largest] {
		largest = left
	}

	if right < n && items[right] > items[largest] {
		largest = right
	}

	if largest != index {
		items[index], items[largest] = items[largest], items[index]
		heapify(items, n, largest)
	}

}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
