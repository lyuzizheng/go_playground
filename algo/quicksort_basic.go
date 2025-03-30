package algo

func QucikSortNew(input []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	var pivot = input[startIndex]
	var mark = startIndex
	// var mark2 = startIndex
	// Swap
	for i := startIndex + 1; i <= endIndex; i++ {
		if input[i] < pivot {
			mark++
			input[mark], input[i] = input[i], input[mark]
		}
	}

	input[startIndex], input[mark] = input[mark], input[startIndex]

	QucikSortNew(input, startIndex, mark-1)
	QucikSortNew(input, mark+1, endIndex)

}

func QucikMergeSort(input []int) []int {

	if len(input) <= 1 {
		return input
	}
	// Split
	var middle = len(input) / 2
	var left = QucikMergeSort(input[:middle])
	var right = QucikMergeSort(input[middle:])

	// Merge
	var result = make([]int, 0, len(left)+len(right))
	var l, r = 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	// Return
	return result
}

// func HeapSort(input []int) {
// 	var heapify func(input []int, i int, len int)

// 	heapify = func(input []int, i int, len int) {
// 		largest := i

// 		leftIndex := i*2 + 1
// 		rightIndex := i*2 + 2

// 		if leftIndex < len && input[leftIndex] > input[largest] {
// 			largest = leftIndex
// 		}
// 		if rightIndex < len && input[rightIndex] > input[largest] {
// 			largest = rightIndex
// 		}
// 		// children bigger
// 		if largest != i {
// 			input[i], input[largest] = input[largest], input[i]
// 			heapify(input, largest, len)
// 		}
// 	}

// 	// Build heap
// 	for i := len(input)/2 - 1; i >= 0; i-- {
// 		heapify(input, i, len(input))
// 	}

// 	fmt.Println(input)

// 	// Sort
// 	for i := len(input) - 1; i >= 0; i-- {
// 		input[i], input[0] = input[0], input[i]
// 		heapify(input, 0, i)
// 	}
// }

func HeapSort(input []int) {

	var heapify func(input []int, i int, len int)
	heapify = func(input []int, i int, len int) {
		largestIndex := i

		leftChild := i*2 + 1
		rightChild := i*2 + 2

		if leftChild < len && input[leftChild] > input[largestIndex] {
			largestIndex = leftChild
		}

		if rightChild < len && input[rightChild] > input[largestIndex] {
			largestIndex = rightChild
		}

		if largestIndex != i {
			input[largestIndex], input[i] = input[largestIndex], input[i]
			heapify(input, largestIndex, len)
		}
	}

	// Build heap
	for i := len(input)/2 - 1; i >= 0; i-- {
		heapify(input, i, len(input))
	}

	// sort

	for i := len(input) - 1; i >= 0; i-- {
		input[i], input[0] = input[0], input[i]
		heapify(input, 0, i)
	}

}
