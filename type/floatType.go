package typeAll

import (
	"fmt"
	"math"
)

func Test5() {
	// 当整数大于23bit能表达的范围,float32的表示将出现误差
	var f float32 = 16777216
	fmt.Println(f == f+1)

	f = 1.903218839298
	fmt.Printf("%g\n", f)

	var z float64
	// 打印出来的结果为0 -0 +Inf -Inf NaN
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	a := math.IsNaN(67)
	b := math.IsNaN(math.NaN())
	fmt.Println(a, b)
}
