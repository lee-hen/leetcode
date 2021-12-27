package fruit_into_baskets

func TotalFruit(fruits []int) int {
	var amount, curr, cnt, a, b int

	for _, kind := range fruits {
		if kind == a || kind == b {
			curr++
		} else {
			curr = cnt+1
		}

		if kind == b {
			cnt++
		} else {
			cnt = 1
			a, b = b, kind
		}

		amount = Max(amount, curr)
	}

	return amount
}

type fruit struct {
	kind int
	amount int
}

func totalFruit(arr []int) int {
	fruits := createFruits(arr)

	if len(fruits) == 1 {
		return fruits[0].amount
	}

	bucket1 := make(map[int]int)
	bucket2 := make(map[int]int)

	bucket1[fruits[0].kind] = fruits[0].amount
	bucket2[fruits[1].kind] = fruits[1].amount
	var total int
	for i := 2; i < len(fruits); i++ {
		if len(bucket1) == 0 {
			bucket1[fruits[i].kind] = fruits[i].amount
			continue
		}

		if len(bucket2) == 0 {
			bucket2[fruits[i].kind] = fruits[i].amount
			continue
		}

		_, ok1 := bucket1[fruits[i].kind]
		_, ok2 := bucket2[fruits[i].kind]

		if !ok1 && !ok2  {
			total = Max(total, getBucketAmount(bucket1) + getBucketAmount(bucket2))

			if _, ok := bucket1[fruits[i-1].kind]; ok {
				bucket2 = make(map[int]int)
				bucket2[fruits[i].kind] = fruits[i].amount
				bucket1[fruits[i-1].kind] = fruits[i-1].amount
			} else if _, ok = bucket2[fruits[i-1].kind]; ok {
				bucket1 = make(map[int]int)
				bucket1[fruits[i].kind] = fruits[i].amount
				bucket2[fruits[i-1].kind] = fruits[i-1].amount
			}
		} else {
			if ok1 {
				bucket1[fruits[i].kind] += fruits[i].amount
			} else if ok2 {
				bucket2[fruits[i].kind] += fruits[i].amount
			}
		}
	}
	total = Max(total, getBucketAmount(bucket1) + getBucketAmount(bucket2))
	return total
}

func createFruits(arr []int) []fruit {
	fruits := make([]fruit, 0)
	fruits = append(fruits, fruit{arr[0], 1})

	for i := 1; i < len(arr); i++ {
		if fruits[len(fruits)-1].kind == arr[i] {
			fruits[len(fruits)-1].amount++
		} else {
			fruits = append(fruits, fruit{arr[i], 1})
		}
	}

	return fruits
}

func getBucketAmount(bucket map[int]int) int {
	for _, num := range bucket {
		return num
	}

	return -1
}

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}
