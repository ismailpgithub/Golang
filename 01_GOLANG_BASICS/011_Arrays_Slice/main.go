package main

import "fmt"

func main(){
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a:= names[0:2]
	b:= names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[2:5]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[3:]
	printSlice(s)

	a1 := make([]int, 5)
	printSlice(a1)

	b1 := make([]int, 0, 10)
	printSlice(b1)

	c1 := b1[:2]
	printSlice(c1)

	d1 := c1[2:5]
	printSlice(d1)

	d1 = append(d1, 1, 2, 3)
	printSlice(d1)

	d1 = append(d1, 4)
	printSlice(d1)

}

func printSlice(s []int){
	fmt.Printf("len=%d  cap=%d  %v\n", len(s), cap(s), s)
}
