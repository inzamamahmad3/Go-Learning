package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hurray! My very first go program!")
	s := "Inzamam"
	fmt.Println(s);
	var i int = 42
	var b string = strconv.Itoa(i)
	fmt.Printf("%q", b)

	// format specifier

	// Constant in Golang Practice

	// const <name> <type> = <value>
	const n string = "Harry Potter"
	const z = "Dumber"

}