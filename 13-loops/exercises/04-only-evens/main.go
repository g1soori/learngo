// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {

	msg := `provide two numbers
	go run 10 20`

	if len(os.Args) != 3 {
		fmt.Println(msg)
		return
	}

	min, errMin := strconv.Atoi(os.Args[1])
	max, errMax := strconv.Atoi(os.Args[2])

	if errMin != nil || errMax != nil {
		fmt.Println(msg)
		return
	}

	if min >= max {
		fmt.Printf("Max value should be greater than min\n%s", msg)
		return
	}

	var sum int

	for i := min; i <= max; i++ {

		if i%2 == 0 {
			sum += i

			fmt.Printf("%d ", i)

			if i != max {
				fmt.Printf("+ ")
			}
		}

	}

	fmt.Printf("= %d\n", sum)

}
