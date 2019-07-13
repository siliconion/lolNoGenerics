package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("lol no generics")
	fmt.Println("import with math/rand, rand.Intn(100): ", rand.Intn(100))
	fmt.Println("import with math, capitals are exported, math.Pi: ", math.Pi)
	fmt.Println(swap(1, 2))
	fmt.Println(split(12345))

	var i int
	fmt.Println("var i int if not assigned is zero value: ", i)
	var j, k int = 1, 2
	fmt.Println("var j, k int = 1, 2, values are assigned: ", j, k)
	l, m, n := 3, true, "lol"
	fmt.Println("use := short assignment then you don't need var and type: ", l, m, n)
	fmt.Printf("l val: %v, type %T\n", l, l)
	fmt.Printf("m val: %v, type %T\n", m, m)
	fmt.Printf("n val: %v, type %T\n", n, n)
	fmt.Println("type() will change types, string(int): ", string(l))
	c := 10 + 20i
	fmt.Println("Complex number! ", c)
	const constant = 1 << 10
	fmt.Println("Constant declare with const: ", constant)
	fmt.Println(" About control flows")
}

// function is
// func name (input) output {}
// can also be
// func swap(x int, y int) (int, int) {
func swap(x, y int) (int, int) {
	return y, x
}

// can specify return var. can "naked return"
func split(sum int) (dec, rmd int) {
	dec = sum / 10
	rmd = sum % 10
	return
}
