
package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i:=2; i<=max;i++{
		var isPrime bool=true
		for j:=2;j*j<i;j++{
			if 0==i%j {
				isPrime=false
				break
			}
		}
		if isPrime==true {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
