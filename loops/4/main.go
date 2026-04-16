
package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		output := ""

		if i%3 == 0 {
			output += "fizz"
		}
		if i%5 == 0 {
			output += "buzz"
		}

		if output == "" {
			fmt.Println(i)
		} else {
			fmt.Println(output)
		}
}}

// don't touch below this line

func main() {
	fizzbuzz()
}
