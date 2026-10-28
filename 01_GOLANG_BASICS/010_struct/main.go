package main

import "fmt"

type Vertex struct{
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

func main(){

	v := Vertex{1, 2}
	p := &v
	p.X = 1e8

	fmt.Println(p.X)

	fmt.Println(v1, v2, v3, v4)
}
