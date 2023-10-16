package daily_contest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	var root1Slice = make([]int, 0)
	var root2Slice = make([]int, 0)

	inOrder(root1, &root1Slice)
	inOrder(root2, &root2Slice)

	return merge(&root1Slice, &root2Slice)

}

func inOrder(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, slice)
	*slice = append(*slice, root.Val)
	inOrder(root.Right, slice)
}

// merge two sorted array
func merge(slice1 *[]int, slice2 *[]int) []int {
	var result = make([]int, 0)
	var i = 0
	var j = 0
	for i < len(*slice1) && j < len(*slice2) {
		if (*slice1)[i] < (*slice2)[j] {
			result = append(result, (*slice1)[i])
			i++
		} else {
			result = append(result, (*slice2)[j])
			j++
		}
	}
	for i < len(*slice1) {
		result = append(result, (*slice1)[i])
		i++
	}
	for j < len(*slice2) {
		result = append(result, (*slice2)[j])
		j++
	}
	return result
}
