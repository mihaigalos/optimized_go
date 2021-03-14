package main

/*
#cgo LDFLAGS: -L./lib -ldotproduct
#include "./lib/interface.h"
*/
import "C"
import (
	"fmt"
	"math/rand"
	"time"
)

func randomSlice() []int32 {
	var result []int32
	for i := 0; i < 350000000; i++ {
		result = append(result, rand.Int31n(100000))
	}
	return result
}

func main() {
	t1 := time.Now()
	s1 := randomSlice()
	s2 := randomSlice()
	t2 := time.Now()
	fmt.Print(t2.Sub(t1), ", ")

	t3 := time.Now()
	result := make([]int32, len(s1), len(s1))
	C.dot_product((*C.int)(&s1[0]), (*C.int)(&s2[0]), C.int(len(s1)), (*C.int)(&result[0]))
	fmt.Print(t3.Sub(t2), ", ")
	fmt.Println("Go+Rust")
}
