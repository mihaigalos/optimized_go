/*
i5-4570T [1549.246, 3600,0000] Mhz

Sequencial, 50000000 elements, dot product
generating random maps..
19.572541915s
starting..
15.856622657s

Parallel, 4 cores, 50000000 elements, dot product
generating random maps..
19.614979378s
starting..
parallelizing..
6.633762924s

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Parallelize(functions ...func()) {
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

func randomMap() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 50000000; i++ {
		m[i] = rand.Intn(100000)
	}
	return m
}
func parallelProduct(m1 *map[int]int, m2 *map[int]int, start int, stop int) map[int]int {
	result := make(map[int]int)
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

	result1 := make(map[int]int)
	result2 := make(map[int]int)
	result3 := make(map[int]int)
	result4 := make(map[int]int)

	fmt.Println("starting..")
	func1 := func() {
		result1 = parallelProduct(&m1, &m2, 0, len(m1)/4)
	}
	func2 := func() {
		result2 = parallelProduct(&m1, &m2, len(m1)/4+1, len(m1)/2)
	}
	func3 := func() {
		result3 = parallelProduct(&m1, &m2, len(m1)/2+1, len(m1)/2+len(m1)/4)
	}
	func4 := func() {
		result4 = parallelProduct(&m1, &m2, len(m1)/2+len(m1)/4+1, len(m1))
	}
	fmt.Println("parallelizing..")
	Parallelize(func1, func2, func3, func4)
	t3 := time.Now()
	fmt.Println(t3.Sub(t2))

	fmt.Println(len(result1))
	fmt.Println(len(result2))
	fmt.Println(len(result3))
	fmt.Println(len(result4))

}
