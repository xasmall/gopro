package typeAll

import "fmt"

func Test9() {
	s := "hello world"
	fmt.Println(len(s))

	l := len(s)
	// 字符串操作s[i:j]基于原始s字符串的第i个字节到第j个字节（不包括j）生成一个新的字符串，
	subS := s[0 : l/2]
	// 如果索引超出字符串范围或者j小于i的话将会导致panic异常
	fmt.Println(subS)
	// 不管是i或者是j都可以被忽略,当它们被忽略的时候，将会采用0作为开始位置，len(s)作为结束位置
	fmt.Println(s[:3])
	fmt.Println(s[3:])
	fmt.Println(s[:])

	// +操作符将两个字符串连接构成一个新的字符串
	fmt.Println("goodbye" + s)

	v := "hello"
	//字符串可以使用==和< >进行比较，比较通过逐个字符比较完成
	t := s < v
	fmt.Println(t)

	// 字符串的值是不可变的，一个字符串包含的字节序列永远不会被改变，但是我们可以给一个字符串变量分配一个新的字符串值
	s += v
	fmt.Println(s)
	s = "hah"
	fmt.Println(s)
	// 因为字符串是不可修改的，尝试修改字符串内部数据的操作也是禁止的
	c1 := s[0]
	// s[0] = 't' compile error
	fmt.Printf("%c\n", c1)

	b1 := []byte(s)
	s = string(b1)
	fmt.Println(b1)
	fmt.Println(s)

}
func Basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
