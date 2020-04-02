package letcode

import (
	"code.qschou.com/peduli/go_common/utils"
	"fmt"
	"strings"
	"testing"
)

func TestTrap(t *testing.T) {
	fmt.Println(Trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(1.0==1.00)
	fmt.Println(utils.Round(1.00,2))
	fmt.Println(strings.Join([]string{},","))
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}