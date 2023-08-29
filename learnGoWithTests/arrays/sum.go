package arrays

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	sums := []int{}

	for _, arr := range arrays {
		sums = append(sums, Sum(arr))
	}
	return sums
}

func SumAllTails(arrays ...[]int) []int {
	sums := []int{}

	for _, arr := range arrays {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			arr = arr[1:]
			sums = append(sums, Sum(arr))
		}
	}
	return sums
}
