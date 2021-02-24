package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

// 其他定义映射的方式：若顶级类型只是一个类型名，你可以在文法的元素中省略它。
var mapTest = map[string]Vertex{
	"Luodayu":   {40.23434, -94.93453},
	"luoweilin": {20.34532, -94.34522},
}

func main() {
	m := make(map[string]Vertex)
	m["Bell lab"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell lab"])

	// 插入元素
	mapTest["liyange"] = Vertex{20.33243, 83.95839}
	fmt.Println("Insert element", mapTest)

	// 修改元素
	mapTest["Luodayu"] = Vertex{12.12345, -12.12345}
	fmt.Println("Modify element", mapTest)

	// 删除元素
	delete(mapTest, "Luodayu")
	fmt.Println("Delete element", mapTest)

	// 通过双赋值检测某个键是否存在
	elem, ok := mapTest["luoweilin"]
	fmt.Println("The value:", elem, "present? ", ok)

}
