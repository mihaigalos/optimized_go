package main

import "math/rand"

func randomMap() map[int]uint32 {
	m := make(map[int]uint32)
	for i := 0; i < 50000000; i++ {
		m[i] = rand.Uint32(100000)
	}
	return m
}

func main() {

}
