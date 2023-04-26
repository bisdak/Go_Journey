package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Shadowing Variables
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	//if statement
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	if m := rand.Intn(10); m == 0 {
		fmt.Println("That's too low")
	} else if m > 5 {
		fmt.Println("That's too big:", m)
	} else {
		fmt.Println("That's a good number:", m)
	}

	//for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}

	//switch statement
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}
