package tourOfGo

import (
	"fmt"
	"math"
)

// Only those variables can be exported which start with a capital letter
// Here math has Pi as exported variable and not pi
func CheckExported() {
	fmt.Println(math.Pi)
}
