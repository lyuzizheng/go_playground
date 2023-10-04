package collections

type TreeNode struct {
	val      interface{}
	children []*TreeNode
}

func NewTree(val interface{}) *TreeNode {
	return &TreeNode{
		val:      val,
		children: []*TreeNode{},
	}
}
