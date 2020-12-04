package stack

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	//n3 := TreeNode{
	//    Val:   3,
	//    Left:  nil,
	//    Right: nil,
	//}
	//n2 := TreeNode{
	//    Val:   2,
	//    Left:  &n3,
	//    Right: nil,
	//}
	//n1 := TreeNode{
	//    Val:   1,
	//    Left:  nil,
	//    Right: nil,
	//}

	ints := inorderTraversal(nil)
	fmt.Printf("%v", ints)
}
