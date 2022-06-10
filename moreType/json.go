package moreType

import (
	"encoding/json"
	"fmt"
	"log"
)

type Actcor struct {
	Name    string `json:"name"`
	Age     int
	Address string
}

type Movie struct {
	Title string
	//omitempty表示当结构体成员为空或者零值的时候不生成该json对象
	Year   int `json:"year,omitempty"`
	Actors Actcor
}

func Test7() {
	var movies = []Movie{{Title: "hellokitty", Year: 2021, Actors: Actcor{Name: "张三", Age: 18, Address: "成都"}},
		{Title: "name", Year: 2021, Actors: Actcor{
			Name: "张三", Age: 18, Address: "马尔康"}}}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json Marshal failed: %s", err)
	}
	fmt.Println(movies)
	// go中结构体中的变量也需要大写，大写为公有
	fmt.Println(string(data))
	// 后面两个参数分别表示每一行输出的前缀和每一个层级的缩进
	// data, err = json.MarshalIndent(movies, "--", " ")
	// if err != nil {
	// 	log.Fatalf("JSon MarshalIndent failed: %s", err)
	// }
	// fmt.Println(string(data))

	// 解码操作 可以定义合适的结构体，选择性的解码json中感兴趣的成员
	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSon unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
