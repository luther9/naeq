package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Returns the square root of n. We add one to the return value to be safe.
func sqrt(n int) int {
	return int(math.Sqrt(float64(n))) + 1
}

// An object that can test if certain numbers are prime. The methods of this
// type are meant to encapsulate the index, which is internally off by 2.
type primeTester []bool

// Sets the maximum number (exclusive) that can be tested for primality.
func (self *primeTester) setMax(max int) {
	// Set all the new numbers to true
	oldMax := len(*self) + 2
	for i := oldMax; i < max; i++ {
		*self = append(*self, true)
	}

	for i := 2; i < sqrt(max); i++ {
		if self.prime(i) {
			// i is a known prime. Make all its multiples false.
			for j := i * i; j < max; j += i {
				(*self)[j-2] = false
			}
		}
	}
}

// Returns true iff n is prime. n must be less than the number set by setMax.
func (self primeTester) prime(n int) bool {
	return self[n-2]
}

func main() {
	// Convert the arguments to a single string. Use a " " seperator to ensure
	// numbers don't run together.
	phrase := strings.ToLower(strings.Join(os.Args[1:], " "))

	value := 0
	// Count each rune
	for _, c := range phrase {
		if unicode.IsLower(c) {
			value += int(c-'a')*19%26 + 1
		}
	}
	// Add in the numbers
	nSlice := regexp.MustCompile(`\d+`).FindAllString(phrase, -1)
	for _, s := range nSlice {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Errorf("The regex caught a non-number: %s", s))
		}
		value += n
	}

	primeFactors := []int{}
	if value > 1 {
		// Get list of primes
		maxPrime := sqrt(value)
		pt := &primeTester{}
		pt.setMax(maxPrime)
		primes := []int{}
		for i := 2; i < maxPrime; i++ {
			if pt.prime(i) {
				primes = append(primes, i)
			}
		}

		// Factor the value
		n := value
		for _, p := range primes {
			if p*p > n {
				break
			}
			for n%p == 0 {
				primeFactors = append(primeFactors, p)
				n /= p
			}
		}
		if n > 1 {
			primeFactors = append(primeFactors, n)
		}
	}

	fmt.Printf("%d %v\n", value, primeFactors)
}
