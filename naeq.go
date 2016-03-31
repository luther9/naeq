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

// Returns the square root of n.
func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// Stores prime numbers.
type primeList struct {
	max   int
	sieve map[int]bool
	list  []int
}

// Sets the maximum possible prime number.
func (pl *primeList) setMax(max int) {
	if max >= pl.max {
		max++
		// Set all the new numbers to true
		oldMax := pl.max
		pl.max = max
		for i := oldMax; i < max; i++ {
			pl.sieve[i] = true
		}

		for i := 2; i <= sqrt(max); i++ {
			if pl.sieve[i] {
				// i is a known prime. Delete all its multiples.
				for j := i * i; j <= max; j += i {
					delete(pl.sieve, j)
				}
			}
		}

		for i := oldMax; i <= max; i++ {
			if pl.sieve[i] {
				pl.list = append(pl.list, i)
			}
		}
	}
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
		primes := &primeList{max: 2, sieve: map[int]bool{}}
		primes.setMax(maxPrime)

		// Factor the value
		n := value
		for _, p := range primes.list {
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
