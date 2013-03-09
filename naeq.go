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

// The constructor and method of this type are meant to encapsulate the off-by-2
// issue.
type primeTester []bool

// Returns a primeTester that can test the primality of numbers in the interval
// [2:length].
func newPrimeTester(length int) primeTester {
	pt := make(primeTester, length-2)
	for i := range pt {
		pt[i] = true
	}
	for i := 2; i < sqrt(length); i++ {
		if pt.prime(i) {
			for j := i * i; j < length; j += i {
				pt[j-2] = false
			}
		}
	}
	return pt
}

// Returns true iff n is prime. n must be within the range set by
// newPrimeTester.
func (self primeTester) prime(n int) bool {
	return self[n-2]
}

func main() {
	// Convert the arguments to a single string
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
		possiblePrimes := newPrimeTester(maxPrime)
		primes := []int{}
		for i := 2; i < maxPrime; i++ {
			if possiblePrimes.prime(i) {
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
