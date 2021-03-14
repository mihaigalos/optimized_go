package main

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
func parallelProduct(m1 *[]int32, m2 *[]int32, start uint32, stop uint32) []int32 {
	result := make([]int32, len(*m1))
	for i := start; i < stop; i++ {
		result[i] = (*m1)[i] * (*m2)[i]
	}
	return result
}

func main() {
	t1 := time.Now()
	//fmt.Println("Generation of random slices, Dotproduct, Algo")
	m1 := randomSlice()
	m2 := randomSlice()

	result1 := make([]int32, len(m1), len(m1))
	result2 := make([]int32, len(m1), len(m1))
	result3 := make([]int32, len(m1), len(m1))
	result4 := make([]int32, len(m1), len(m1))

	func1 := func() {
		result1 = parallelProduct(&m1, &m2, 0, uint32(len(m1)/4))
	}
	func2 := func() {
		result2 = parallelProduct(&m1, &m2, uint32(len(m1)/4), uint32(len(m1)/2))
	}
	func3 := func() {
		result3 = parallelProduct(&m1, &m2, uint32(len(m1)/2), uint32(len(m1)/2+len(m1)/4))
	}
	func4 := func() {
		result4 = parallelProduct(&m1, &m2, uint32(len(m1)/2+len(m1)/4), uint32(len(m1)))
	}

	t2 := time.Now()
	parallelize(func1, func2, func3, func4)
	t3 := time.Now()
	fmt.Print(t2.Sub(t1), ", ")
	fmt.Print(t3.Sub(t2), ", ")
	fmt.Println("PureGo")
}
