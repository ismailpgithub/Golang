package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Pinjari"
	lastName := "Ismail"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))

}
