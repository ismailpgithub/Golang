package main

import (
	"fmt"
	"math"
)

type Abser interface{
	Abs() float64
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


func main(){
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// f, k := i.(float64)
	// fmt.Println(f, k)
	x := Person{"Ravi", 42}
	y := Person{"Raj", 23}

	fmt.Println(x, y)

}


type MyFloat float64

func (f MyFloat) Abs() float64{
	if f < 0{
		return  float64(-f)
	}

	return float64(f)
}

type Vertex struct{
	X, Y float64
}

func (v *Vertex) Abs() float64{
	return  math.Sqrt(v.X * v.X + v.Y * v.Y)
}

