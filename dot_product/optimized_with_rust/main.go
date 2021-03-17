package main

/*
#cgo LDFLAGS: -L./lib -ldotproduct
#include "./lib/interface.h"
*/
import "C"
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

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

	result1 := make([]int32, len(s1), len(s1))
	result2 := make([]int32, len(s1), len(s1))
	result3 := make([]int32, len(s1), len(s1))
	result4 := make([]int32, len(s1), len(s1))

	func1 := func() {
		C.dot_product((*C.int)(&s1[0]), (*C.int)(&s2[0]), 0, C.ulonglong(len(s1)/4), (*C.int)(&result1[0]))
	}
	func2 := func() {
		C.dot_product((*C.int)(&s1[0]), (*C.int)(&s2[0]), C.ulonglong(len(s1)/4), C.ulonglong(len(s1)/2), (*C.int)(&result2[0]))
	}
	func3 := func() {
		C.dot_product((*C.int)(&s1[0]), (*C.int)(&s2[0]), C.ulonglong(len(s1)/2), C.ulonglong(len(s1)/2+len(s1)/4), (*C.int)(&result3[0]))
	}
	func4 := func() {
		C.dot_product((*C.int)(&s1[0]), (*C.int)(&s2[0]), C.ulonglong(len(s1)/2+len(s1)/4), C.ulonglong(len(s1)), (*C.int)(&result4[0]))
	}

	t3 := time.Now()
	parallelize(func1, func2, func3, func4)
	t4 := time.Now()
	fmt.Print(t4.Sub(t3), ", ")
	fmt.Println("Go+Rust")
}
