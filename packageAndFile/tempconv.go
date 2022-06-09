package tempconv

// 类型

// 一个类型声明语句创建了一个新的类型名称,和现有的类型具有相同的底层结构,但是它们是不同的类型,不可以互相比较或者混在一个表达式中运算

// 对于一个自定义的包来说，第一个字母小写是private，第一个字母大写是public

type Celsius float64
type Fahreheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahreheit {
	return Fahreheit(c*9.5 + 32)
}

func FToC(f Fahreheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
