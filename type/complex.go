package typeAll

import (
	"fmt"
	"math/cmplx"
)

// go语言提供了两种精度的复数类型，complex64和complex128，分别对应的是float32和float64

func Test6() {
	var x complex64 = complex(6, 4) //6+4i
	var y complex64 = complex(1, 3) // 1 + 3i
	fmt.Println(x + y)

	// real()获取实数部
	fmt.Println(real(x))
	// imag()获取虚部
	fmt.Println(imag(x))

	//如果一个浮点数面值或者一个十进制数面值后面跟着一个i，它将构成一个复数的虚部
	fmt.Println(1i * 1i) //结果应该是-1

	// math的cmplx提供了很多复数处理的函数，下面可以求复数的平方根函数和求幂函数
	fmt.Println(cmplx.Sqrt(-1))
	fmt.Println(cmplx.Pow(complex128(x), complex128(y)))
}
