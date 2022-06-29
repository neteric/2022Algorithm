package linklist

import (
	"fmt"
	"testing"
)

func TestMakeLinkListBySlice(*testing.T) {

	l := MakeLinkListBySlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(traverseList(l))
}
