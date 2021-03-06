package main

import (
	"flag"
	"fmt"
	"strconv"
)

func isProth(number int) bool {
	// A Proth number is a number N of the form N=k(2^n)+1

	// deduct 1 from the number to give a number of the form k(2^n)
	n := number - 1
	// loop through all odd numbers from k = 1 to n/k
	// check if k can divide n in a way that n/k is a power of 2
	for k := 1; k < (n / k); k += 2 {
		// Check if n is divisible by k
		if n%k == 0 {
			// Check if n/k is a power of two
			if isPowerOfTwo(n / k) {
				return true
			}
		}
	}
	return false
}

func isPowerOfTwo(n int) bool {
	return (n != 0) && ((n & (n - 1)) == 0)
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isProthPrime(n int) bool {
	// if n is a proth number then check if it is a prime number
	if isProth(n) {
		fmt.Printf("%v is a Proth number, \n", n)
		if isPrime(n) {
			return true
		}
		return false
	}
	fmt.Printf("%v is not a Proth number, \n", n)
	return false
}

func main() {
	flag.Parse()
	args := flag.Args()
	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}
	v := isProthPrime(n)
	if v {
		fmt.Printf("%v is a Proth prime \n", n)
	} else {
		fmt.Printf("%v is not a Proth prime \n", n)
	}
}
