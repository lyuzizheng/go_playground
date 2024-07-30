package daily_contest

import "fmt"



func CountPairs(root *TreeNode, distance int) int {

    var result = 0

    var findChildren func(root *TreeNode, distance int) (distances []int, curResult int)
    

    findChildren = func(root *TreeNode, distance int) (distances []int, curResult int){
        if root == nil {
            return make([]int,0), 0
        }
        var result = 0
        leftChild, leftResult := findChildren(root.Left, distance)
        rightChild, rightResult := findChildren(root.Left, distance)
        result = result + leftResult + rightResult

        for leftDistance := range leftChild {
            for rightDistance := range rightChild {
                if leftDistance + rightDistance <= distance {
                    result ++
                }
            } 
        }

		fmt.Println(result, leftChild, rightChild, leftResult, rightResult)

        newDistances := make([]int,0)
        for leftDistance := range leftChild {
            if leftDistance + 1 < distance {
                newDistances = append(newDistances, leftDistance+1)
            }
        }
        for rightDistance := range rightChild {
            if rightDistance + 1 < distance {
                newDistances = append(newDistances, rightDistance+1)
            }
        }

        return newDistances, result

    }

    _, result = findChildren(root, distance)




    return result
}
