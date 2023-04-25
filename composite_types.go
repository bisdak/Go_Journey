package main

import "fmt"

func main() {
	//Arrays
	//arrays are rarely used in Go

	var x [3]int
	fmt.Println(x)
	var y = [3]int{10,20,30}
	fmt.Println(y)

	var z = [...]int{10,20,30}
	fmt.Println(z)
	fmt.Println(y==z)
	fmt.Println(len(z))

	//Slices
	var _x = []int{10,20,30}
	fmt.Println(_x)
	fmt.Println(_x[0])

	var _y []int
	fmt.Println(_y)
	fmt.Println(_y==nil)

	//append
	var j []int
	var k = []int{50, 60, 70}
	j = append(j, 10)
	j = append(j, 20, 30, 40)
	j = append(j, k...)
	fmt.Println(j)

	// capacity
	fmt.Println(j, len(j), cap(j))

	// make
	// create slice initial capacity
	l := make([]int, 5)
	fmt.Println(l, len(l), cap(l))
	l = append(l, 10)
	fmt.Println(l, len(l), cap(l))
	a := make([]int, 5, 10)
	fmt.Println(a, len(a), cap(a))

	//Slicing Slices
	b := []int{1,2,3,4}
	c := b[:2]
	e := b[1:]
	f := b[:]
	fmt.Println(b,c,e,f)

	//Converting Arrays to Slices
	//Same to Slice

	//Copy
	g := make([]int, 5)
	num := copy(g, b)
	fmt.Println(g, num)

	//Strings and Runes and Bytes
	var s string = "Hello there"
	var by byte = s[6]
	var st string = s[6:10]
	fmt.Println(s, by, st)

	//Maps
	teams := map[string][]string{
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(teams["Orcas"])
	// age := make(map[int][]string, 10)

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	_v, _ok := m["goodbye"]
	fmt.Println(_v, _ok)
}