package lib

func RandomArray(N, min, max int) []int {
	array := make([]int, N)

	for j := 0; j < N; j++ {
		array[j] = RandomIntInRange(min, max)
	}

	return array
}
