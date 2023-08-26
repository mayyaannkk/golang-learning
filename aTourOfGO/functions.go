package tourOfGo

import "fmt"

// Defining a function in Go
func concat(x string, y string) string {
	return x + y
}

// function arguments can specify the data type only one if they share the same data-type. But it has to start from right
func add(x, y int) int {
	return x + y
}

// Multiple returns
func swap(x, y string) (string, string) {
	return y, x
}

// Naked returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(concat("abc", "def"))
	fmt.Println(add(42, 13))
}
