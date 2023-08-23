package exported

import (
	"fmt"
	"math"
)

// Only those variables can be exported which start with a capital letter
// Here math has Pi as exported variable and not pi
func main() {
	fmt.Println(math.Pi)
}
