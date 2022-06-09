package typeAll

import "fmt"

func Test2() {
	var u uint8 = 234
	fmt.Println(u)

	var i int8 = 127
	// 运算结果溢出
	fmt.Println(i, i+1, i*i)
}

func Test3() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Println(x, y)
}
