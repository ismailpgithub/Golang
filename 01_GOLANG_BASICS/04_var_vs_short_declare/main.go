package main

import (
	"fmt"
)

func main() {
	var city string = "London"
	var country = "france"

	subcribers := 5000
	subcribers += 1000

	like, comments := 1000, "30"

	fmt.Println(city, country, subcribers, like, comments)
}
