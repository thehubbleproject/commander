package core

import (
	"fmt"
	"testing"
)

func TestGetAdjacentNodePath(t *testing.T) {
	res, err := GetAdjacentNodePath("00000001")
	fmt.Println("data", res, err)
}

func TestGetParentPath(t *testing.T) {
	leftChildPath := "000"
	// rightChildPath := "111"
	// expectedParentPath := "11"
	fmt.Println(GetParentPath(leftChildPath))
}
