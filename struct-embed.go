package main

import "fmt"

//Go supports embedding of structs and interfaces
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base //embedding looks like a field without a name
	str  string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "Aiyan",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("Num:", co.base.num)

	fmt.Println("Describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("Describer:", d.describe())
}
