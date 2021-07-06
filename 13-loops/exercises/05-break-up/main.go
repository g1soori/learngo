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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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

	for {

		if min%2 == 1 {
			min++
			continue
		}

		sum += min

		fmt.Printf("%d ", min)

		if min != max {
			fmt.Printf("+ ")
		} else {
			break
		}

		min++

	}

	fmt.Printf("= %d\n", sum)

}
