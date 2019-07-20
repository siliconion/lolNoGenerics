package main

import "fmt"

func main() {
	pointers()
	structs()
	slices()
	sliceLiterals()
	makeSlice()
	maps()
}

func pointers() {
	a := "lol"
	var ap = &a
	fmt.Printf("a is of type %T and value %v, and pointer to a is of type %T, and value %v \n", a, a, ap, ap)
	*ap += " no generics"
	fmt.Println("Assigning through pointers, now we have: ", a)
	fmt.Println("No pointer arithmetic though!")
}

type someStruct struct {
	x int
	y string
}

func structs() {
	fmt.Println("===========================")
	dict := someStruct{y: "lol"}
	dict.y += " no generics" // accessed with dots
	fmt.Println("A struct is a collection of fields. If not assign default to zero value. Like: ", dict)
	np := dict
	np.x = 333
	fmt.Printf("Direct assignment creates a copy. New struct %v, old struct %v \n", np, dict)
	p := &dict
	fmt.Printf("A a pointer to struct is of type %T, and value %v, can access the value with dot: %v \n", p, p, p.x)
	p.x = 1
	fmt.Println("It can change the struct: ", p, dict)
	fmt.Println("It does not change the other assignment?: ", np)

	// how do you get all the keys? loop through this?
}

func slices() {
	fmt.Println("===========================")
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7 ,8, 9, 10}
	fmt.Println("[n]T is an array: ", a)
	fmt.Println("OH NO CANNOT resize!!!!! WTF is this, my junior year? length: ", len(a))

	var s = a[1:8]
	fmt.Println("So there's slices. Which is a flexible representation(?) of an array. You can create with slicing a real array, like in Python:", s)
	fmt.Printf("arrays are of type %T, and slices are of type %T \n", a, s)
	fmt.Println("The original array is unchanged: ", a)
	s[1] = 999
	fmt.Println("But slice can change the underlying array. OMG WHY ", a, s)
	s= a[:5]
	fmt.Println("And you can reslice like...", s)
	s = a[:8]
	fmt.Println("Then like...", s)
	fmt.Printf("The slice %v has length %v and capacity %v \n", a, len(a), cap(a) )

	var snil []int
	fmt.Printf("Unitialized slice %v has length %v and capacity %v \n", snil, len(snil), cap(snil) )
	fmt.Println("And if nil == snil? ", snil == nil)

}

func sliceLiterals()  {
	fmt.Println("===========================")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("slice %v are of type %T \n", s, s)

	// range is basically enumerate in python? is it generator?
	// can use _ to skip. But why bother?
	for i, v := range s{
		fmt.Println("index and val are: ", i, v)
	}
	//s = s[:11] // slice out of bound error
	s = s[:5]
	fmt.Println("slice it :5: ", s)
	fmt.Printf("The slice %v has length %v and capacity %v \n", s, len(s), cap(s) )
	s = s[:9]
	fmt.Println("slice it :9: ", s)
	fmt.Printf("The slice %v has length %v and capacity %v \n", s, len(s), cap(s) )

	s = s[5:8]
	fmt.Println("slice it :9: ", s)
	fmt.Printf("The slice %v has length %v and capacity %v \n", s, len(s), cap(s) )

	//s = s[4:8] now it'll be out of bounds
	//fmt.Printf("The slice %v has length %v and capacity %v \n", s, len(s), cap(s) )

}

func makeSlice()  {
	fmt.Println("===========================")

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	e := append(d, 20)
	printSlice("d", d)
	printSlice("e", e)

	ss := make([][]int, 5)
	fmt.Println(ss)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func maps(){
	fmt.Println("===========================")
	// like structs, but needs key? but struct can have field names(?) too?
	var m = make(map[string]int)
	fmt.Println(m)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m, m["a"])
	
	ml := map[string]someStruct {
		"a": {1, "lol"},
		"b": {3, "no"},
	}
	fmt.Println(ml)
	ml["c"] = someStruct{5, "generics"}
	fmt.Println(ml)
	delete(ml, "a")
	fmt.Println(ml)
	vc, okc := ml["c"]
	fmt.Println(vc, okc)
	va, oka := ml["a"]
	fmt.Println(va, oka)

}