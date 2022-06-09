package main

import (
	"fmt"
	moreType "gopro/moreType"
	typeAll "gopro/type"
)

// 我们注意go导入包的时候，需要在项目的设置go mod文件，这个文件设置了寻址路径

func main() {
	// var c tempconv.Celsius
	// var f tempconv.Fahreheit
	// fmt.Println(c, f)
	// // complie error : type mismatch
	// // fmt.Println(c == f)

	// boilingF := tempconv.CToF(tempconv.BoilingC)
	// // Println打印的是toString()fangfa
	// fmt.Println(boilingF)

	// f = 982
	// // 这样就可以打印它具体的值了
	// fmt.Printf("%g\n", f)
	// // 当前类型一样，变成了可以比较的
	// fmt.Println(boilingF == f)
	// // fmt.Println()

	// s1 := c.String()
	// s2 := f.String()
	// fmt.Println(s1, s2)
	// scope.Test()
	// typeAll.Test2()
	// typeAll.Test3()
	// typeAll.Test5()
	// typeAll.Test6()
	typeAll.Test9()
	bn := typeAll.Basename("hu/hsa.g")
	fmt.Println(bn)
	typeAll.Test10()
	moreType.Test1()
	moreType.Test2()
}
