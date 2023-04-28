package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(p)

	failedUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func stringp(s string) *string {
	return &s
}
