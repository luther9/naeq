package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/bobappleyard/readline"
)

func fatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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

func newPrimeList() *primeList {
	return &primeList{max: 1, sieve: map[int]bool{}}
}

// Sets the maximum possible prime number.
func (pl *primeList) setMax(max int) {
	if max > pl.max {
		// Set all the new numbers to true
		oldMax := pl.max + 1
		pl.max = max
		for i := oldMax; i <= max; i++ {
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

type valueEntry struct {
	value int
	words []string
}

// Returns the numerical value of phrase. Phrase must not have any capital
// letters.
func getValue(phrase string) int {
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
	return value
}

// Outputs a value along with its prime factors.
func outputValue(value int, primes *primeList) {
	primeFactors := []int{}
	if value > 1 {
		primes.setMax(sqrt(value))

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

// Outputs the value of phrase along with its prime factors.
func processPhrase(phrase string, primes *primeList) {
	outputValue(getValue(strings.ToLower(phrase)), primes)
}

func main() {
	primes := newPrimeList()

	fileName := flag.String("f", "", "Output the values of all words found in the specified file.")
	flag.Parse()

	if *fileName != "" {
		text, err := ioutil.ReadFile(*fileName)
		fatal(err)
		words := regexp.MustCompile(`[\w'\-]+`).FindAll(text, -1)
		data := []valueEntry{}
		seenWords := map[string]bool{}
		for _, wordSlice := range words {
			word := strings.ToLower(string(wordSlice))
			if !seenWords[word] {
				seenWords[word] = true
				value := getValue(word)
				i := sort.Search(len(data), func(i int) bool {
					return data[i].value >= value
				})
				if i == len(data) || value != data[i].value {
					data = append(data, valueEntry{})
					copy(data[i+1:], data[i:])
					data[i] = valueEntry{value: value}
				}
				entry := &data[i]
				i = sort.SearchStrings(entry.words, word)
				entry.words = append(entry.words, "")
				copy(entry.words[i+1:], entry.words[i:])
				entry.words[i] = word
			}
		}
		for _, entry := range data {
			outputValue(entry.value, primes)
			for _, word := range entry.words {
				fmt.Println(word)
			}
			fmt.Println()
		}
	} else if len(os.Args) > 1 {
		// Convert the arguments to a single string. Use a " " seperator to ensure
		// numbers don't run together.
		processPhrase(strings.Join(os.Args[1:], " "), primes)
	} else {
		for {
			phrase, err := readline.String("> ")
			if err != nil {
				break
			}
			processPhrase(phrase, primes)
			if phrase != "" {
				readline.AddHistory(phrase)
			}
		}
		fmt.Println()
	}
}
