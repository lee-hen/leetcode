package jewels_and_stones

func NumJewelsInStones(jewels string, stones string) int {
	lookup := make(map[byte]struct{})
	for i := range jewels {
		lookup[jewels[i]] = struct{}{}
	}

	var num int
	for i := range stones {
		if _, ok := lookup[stones[i]]; ok {
			num++
		}
	}
	return num
}
