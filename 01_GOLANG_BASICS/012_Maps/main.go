package main

import "fmt"

type Vertex struct{
	Lat, Long float64
} 

var m1 = map[string]Vertex{
	"Google": Vertex{
		37.4224, -122.0842,
	},
}

func main(){
	m := make(map[string]int)

	m["Answer1"] = 42
	fmt.Println("The value : ", m["Answer1"])

	m["Answer2"] = 48
	fmt.Println("The value : ", m["Answer2"])

	fmt.Println(m)

	m1["Bell Labs"] = Vertex{
		40.68433, -74.3997,
	}

	fmt.Println(m1)

	for key := range m1{
		fmt.Println(key)
	}

}