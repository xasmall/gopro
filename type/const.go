package typeAll

import "fmt"

// 对于批量声明的常量，除了第一个外其他的常量右边的表达式都可以省略，如果省略表达式则表示使用前面常量的初始化表达式写法，对应的常量类型也一样
const (
	a = 1
	b
	c = 2
	d
)

/**
iota常量生成器,它用于生成一组以相似规则初始化的常量，但是不用每行都写一编初始化表达式
在一个const声明语句中，在第一个声明的常量所在的行，iota会被置为0，然后在每个有常量声明的行加1
*/
type Weekday int

const (
	Sunnday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Firday
	Saturday
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

// 注意go语言的常量是无类型常量，无类型常量不仅可以提供更高的运算精度，而且可以直接用于表达式而不需要显式的类型转换

func Test10() {
	fmt.Println(a, b, c, d)
	fmt.Println(Thursday, Saturday)
	// 这里YiB已经任何go语言中整数类型能够表达的范围，但它们仍然是合法的常量，而且下面的表达式依然有效
	fmt.Println(YiB / ZiB)

}
