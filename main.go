package main

import (
	"fmt"

	"github.com/PhanNam1501/Introduction-to-Algorithms/heapsort"
)

func main() {
	h := heapsort.NewMaxHeap([]int{2, 4, 1, 7, 8, 9, 3, 14, 10, 16})
	h.BuildMaxHeap()
	res := h.GetMaxHeap()
	fmt.Println(res)
	h.ExtractMax()
	res = h.GetMaxHeap()
	fmt.Println(res)
}
