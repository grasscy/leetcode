package stack

import (
	"fmt"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {

	n2 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	n4 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	n5 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   20,
		Left:  n4,
		Right: n5,
	}
	n1 := &TreeNode{
		Val:   3,
		Left:  n2,
		Right: n3,
	}

	order := zigzagLevelOrder(n1)
	fmt.Println(order)
}
