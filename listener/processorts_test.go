package listener

import (
	"fmt"
	"math/big"
	"testing"
)

func TestPathConversion(t *testing.T) {
	path := "100101010"
	pathBig := big.NewInt(0)
	pathBig.SetString(path, 2)
	fmt.Println("path", pathBig.String())
	expectedSubPath := "101"
	expectedPathBig := big.NewInt(0)
	expectedPathBig.SetString(expectedSubPath, 2)
	fmt.Println("path", expectedPathBig.String())
	expectedPathBig.
}
