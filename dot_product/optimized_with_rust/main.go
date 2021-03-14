package main

/*
#cgo LDFLAGS: -L./lib -ldotproduct
#include "./lib/interface.h"
*/
import "C"
import (
	"fmt"
	"math/rand"
)

func randomSlice() []int32 {
	var result []int32
	for i := uint32(0); i < 10; i++ {
		result = append(result, rand.Int31n(10000))
	}
	return result
}

func main() {
	s1 := randomSlice()
	s2 := randomSlice()
	result := make([]int32, len(s1), len(s1))
	C.dot_product((*C.int)(&s1[0]), (*C.int)(&s2[0]), C.int(len(s1)), (*C.int)(&result[0]))

	for i, e := range result {
		fmt.Print("[", i, "]: ")
		fmt.Println(s1[i], "*", s2[i], "=", e)
	}
}
