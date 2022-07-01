package linklist

import (
	"fmt"
	"testing"
)

func TestMakeLinkListBySlice(*testing.T) {

	l := MakeLinkListBySlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(traverseList(l))
}

func TestMergeTwoLists(*testing.T) {
	l1 := MakeLinkListBySlice([]int{})
	l2 := MakeLinkListBySlice([]int{0})
	merged := mergeTwoLists(l1, l2)
	fmt.Println(traverseList(merged))
}
