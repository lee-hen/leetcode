package lib

import (
	"math/rand"
	"time"
)

func RandomArray(N, min, max int) []int {
	array := make([]int, N)

	for j := 0; j < N; j++ {
		array[j] = RandomIntInRange(min, max)
	}

	return array
}

func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandomIntInRange(min, max int) int {
	return RandomInt(max+1-min) + min
}
