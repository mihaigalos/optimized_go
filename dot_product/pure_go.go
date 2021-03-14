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

func randomMap() map[uint32]int32 {
	m := make(map[uint32]int32)
	for i := uint32(0); i < 50000000; i++ {
		m[i] = rand.Int31n(100000)
	}
	return m
}
func parallelProduct(m1 *map[uint32]int32, m2 *map[uint32]int32, start uint32, stop uint32) map[uint32]int32 {
	result := make(map[uint32]int32)
	for i := start; i < stop; i++ {
		result[i] = (*m1)[i] * (*m2)[i]
	}
	return result
}

func main() {
	t1 := time.Now()
	fmt.Println("generating random maps..")
	m1 := randomMap()
	m2 := randomMap()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	result1 := make(map[uint32]int32)
	result2 := make(map[uint32]int32)
	result3 := make(map[uint32]int32)
	result4 := make(map[uint32]int32)

	fmt.Println("starting..")
	func1 := func() {
		result1 = parallelProduct(&m1, &m2, 0, uint32(len(m1)/4))
	}
	func2 := func() {
		result2 = parallelProduct(&m1, &m2, uint32(len(m1)/4+1), uint32(len(m1)/2))
	}
	func3 := func() {
		result3 = parallelProduct(&m1, &m2, uint32(len(m1)/2+1), uint32(len(m1)/2+len(m1)/4))
	}
	func4 := func() {
		result4 = parallelProduct(&m1, &m2, uint32(len(m1)/2+len(m1)/4+1), uint32(len(m1)))
	}
	fmt.Println("parallelizing..")
	parallelize(func1, func2, func3, func4)
	t3 := time.Now()
	fmt.Println(t3.Sub(t2))

	fmt.Println(len(result1))
	fmt.Println(len(result2))
	fmt.Println(len(result3))
	fmt.Println(len(result4))

}
