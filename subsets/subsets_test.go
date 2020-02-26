package subsets

import (
	"fmt"
	"testing"
)

func TestRes(t *testing.T) {
	res := subsets([]int{1, 2, 3})
	fmt.Println(res)
}
