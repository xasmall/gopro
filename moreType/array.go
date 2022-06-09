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
}
