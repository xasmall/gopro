package tempconv

import "fmt"

// 注意一个文件夹下面的go文件必须是同一个包

func init() {
	f1 := CToF(BoilingC)

	fmt.Println(f1)
}

// 一个文件可以有多个init函数。init初始化函数除了不能被调用和引用外，其他行为和普通函数类似

func init() {
	f2 := FToC(Fahreheit(FreezingC))
	fmt.Println(f2)
}

// 这个相当于重写了toString()方法
func (c Celsius) String() string {
	return fmt.Sprintf("当前温度是摄氏温度: %g", c)
}

func (f Fahreheit) String() string {
	return fmt.Sprintf("当前温度是华氏温度：%g", f)
}
