package main

import "fmt"

func even(arr []int) []int {
	evenNums := make([]int, 0)
	for _, num := range arr {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}
	return evenNums
}

func odd(arr []int) []int {
	oddNums := make([]int, 0)
	for _, num := range arr {
		if num%2 != 0 {
			oddNums = append(oddNums, num)
		}
	}
	return oddNums
}

func prime(arr []int) []int {
	primeNums := make([]int, 0)
	for _, num := range arr {
		if num < 2 {
			continue
		}
		flag := true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeNums = append(primeNums, num)

		}
	}
	return primeNums
}

func multipleOf5(arr []int) []int {
	multipleNums := make([]int, 0)
	for _, num := range arr {
		if num%5 == 0 {
			multipleNums = append(multipleNums, num)
		}

	}
	return multipleNums
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Array: ", arr)
	fmt.Println("Even: ", even(arr))
	fmt.Println("Odd: ", odd(arr))
	fmt.Println("Prime: ", prime(arr))
	fmt.Println("Odd Prime: ", odd(prime(arr)))
	fmt.Println("Even & Multiple of 5: ", multipleOf5(even(arr)))
}
