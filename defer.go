package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	result := 1
	defer func() {
		result = result * 2 // Deferred function modifies the return value
		fmt.Println("result -- ", result)
	}()

	fmt.Println("Before defer result - ", result)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()")
}

//Deferred function calls are executed in the reverse order they were deferred. The last deferred function is executed first, and so on.
//If a panic occurs in a function, deferred functions in that function will still run before the panic is propagated up the call stack.
// You can use recover within a deferred function to handle or suppress the panic.
