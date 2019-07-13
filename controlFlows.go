package main

import "fmt"

func main() {
	fmt.Println("About control flows")
	controlFlow()
	aboutDefer()
}

func controlFlow() {
	fmt.Println("for loop is for start; end; incr {}")
	fmt.Println("if can have some shit before, if shit; eval {}")
	for i := 0; i < 10; i++ {
		if n := -1; i%2 == 0 {
			fmt.Println(i * n)
		} else {
			fmt.Println(i * 10)
		}
	}
	fmt.Println("or can only have for end{} which is like while.")
	j := 5
	for j < 10 {
		fmt.Println(j)
		j++
	}
}

func aboutDefer() {
	fmt.Println("Defer is a stack")
	for i := 0; i < 10; i++ {
		defer fmt.Println("no generics X", i)
	}

	fmt.Println("lol")
}