package main

import (
	"flag"
	"fmt"
	"strconv"
)

func isProth(in int) bool {
	n := in - 1
	for k := 1; k < (n / k); k += 2 {

		// Check if k can divide n such that n/k is a power of 2
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

func main() {
	flag.Parse()
	args := flag.Args()
	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}
	v := isProth(n)
	if v {
		fmt.Printf("%v is a Proth prime \n", n)
		return
	}
	fmt.Printf("%v is not a proth prime \n", n)
}
