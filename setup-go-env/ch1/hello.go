package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	type firstPerson struct {
		name string
		age  int
	}
	type secondPerson struct {
		name string
		age  int
	}

	f := firstPerson{}
	s := secondPerson{}
	fmt.Println(f == firstPerson(s))
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

}
