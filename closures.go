package main

import "fmt"

func f1() func() int {
	i := 0
	return func() int { //Anonymous functions
		i++
		return i
	}
}

func main() {

	a := f1()

	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3

	b := f1()
	fmt.Println(b()) //1
}
