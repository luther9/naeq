# `naeq`

A [New Aeon English
Qabalah](https://en.wikipedia.org/wiki/English_Qabalah#ALW_Cipher) calculator.

## Usage

`naeq [WORD]...`
`naeq -f FILENAME`

`naeq` takes any string and outputs the sum of the numerical values of its
letters, followed by a list of prime factors of that sum. With no arguments,
`naeq` enters an interactive mode (using readline) where you may repeatedly type
phrases at a prompt and see their results. When passing any non-option
arguments, `naeq` sums them up, outputs the results, and exits. With the `-f`
option, `naeq` will tally up each word in the file given by FILENAME and output
their results in order by value.
