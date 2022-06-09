package moreType

import "fmt"

func Test1() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// index和element
	for i, v := range a {
		fmt.Println(i, v)
	}

	var q []int = []int{1, 3, 4, 6}
	fmt.Println(q[1])

	//	如果出现的是...省略号，则表示数组的长度是根据初始化的值的个数来确定计算的
	p := [...]int{2, 3, 5}
	fmt.Println(len(p))
	// 使用%T可以打印出来参数的类型
	fmt.Printf("%T\n", p)
	// 我们可以使用指定一个索引和对应值列表的方式来进行初始化,这种形式的初始化方式，索引的顺序是无关紧要的
	symbol := [...]string{0: "$", 1: "%", 2: "*"}
	fmt.Println(symbol)
	// 没有的索引可以忽略
	r := [100]int{99: -1}
	fmt.Println(r)
	// 当两个数组的所有元素都是相等的时候数组才是相等的
}
