package algo

import (
	"fmt"

	"briemann.com/wordler/dict"
	"briemann.com/wordler/rules"
)

var (
	ValidGuesses = append(dict.Words, dict.AdditionalGuesses...)
)

func AsciiToIndex(c byte) byte {
	return c - 'a'
}

const (
	CHARS = 26
	WLEN  = 5
)

type CharSet [CHARS]bool
type WordSet [WLEN]CharSet

func RemainingAmount(possibles *WordSet, source []string) int {
	amount := 0

	for w := 0; w < len(source); w++ {
		word := source[w]
		for i := 0; i < 5; i++ {
			char := word[i]
			idx := AsciiToIndex(byte(char))
			exists := possibles[i][idx]
			if !exists {
				goto skip
			}
		}
		amount++
	skip:
	}

	return amount
}

func Remaining(possibles WordSet, source []string) []string {
	remaining := []string{}
	var idx byte

	for w := 0; w < len(source); w++ {
		word := source[w]
		for i := 0; i < 5; i++ {
			char := word[i]
			idx = AsciiToIndex(byte(char))
			if !possibles[i][idx] {
				goto skip
			}
		}
		remaining = append(remaining, word)
	skip:
	}

	return remaining
}

type wordAvg struct {
	Word string
	Avg  int
}

func Guess(possibles *WordSet, remaining *[]string, attempt int, verbose bool) string {
	if verbose {
		fmt.Printf("POSSIBLES: %v\n", possibles)
	}
	// First guess is always "tares".
	if attempt == 1 {
		return "tares"
	}

	rem := Remaining(*possibles, *remaining)
	remaining = &rem
	if verbose {
		fmt.Printf("There are %d possible words: %v\n", len(*remaining), remaining)
	}

	if len(*remaining) == 0 {
		// Should never happen.
		return "ERROR"
	}

	if len(*remaining) <= 2 {
		// If there are only 1 or 2 possible solutions left we can always pick the first.
		return (*remaining)[0]
	}

	// For all possible (guess word, remaining solution) tuples calculate how many solutions
	// would remain after the reduction. Then choose input word by the lowest average.
	averages := make([]wordAvg, len(ValidGuesses))
	possiblesTmp := WordSet{}

	for a := 0; a < len(ValidGuesses); a++ {
		guess := ValidGuesses[a]
		avg := 0
		// TODO: reduce RemainingAmount calls
		for r := 0; r < len(*remaining); r++ {
			secret := (*remaining)[r]
			possiblesTmp = *possibles

			// Calculate hint for (guess, secret) pair.
			hint := rules.Check(guess, secret)

			// Reduce possible remaining words.
			Reduce(&possiblesTmp, guess, hint)
			avg += RemainingAmount(&possiblesTmp, *remaining)
		}
		// Calc the average for the input.
		avg /= len(*remaining)
		averages[a] = wordAvg{
			Word: guess,
			Avg:  avg,
		}
	}

	if verbose {
		fmt.Printf("AVERAGES: %v\n", averages)
	}

	lowest := 1000
	best := ""
	for _, wa := range averages {
		if wa.Avg < lowest {
			best = wa.Word
			lowest = wa.Avg
		}
	}

	return best
}

func Reduce(possibles *WordSet, guess string, hint string) {
	seen := CharSet{}
	for i := 0; i < 5; i++ {
		if hint[i] == 'g' { // green
			for c := 'a'; c <= 'z'; c++ {
				char := byte(c)
				if char != guess[i] {
					possibles[i][AsciiToIndex(char)] = false
				}
			}
			seen[AsciiToIndex(guess[i])] = true
		}
	}
	for i := 0; i < 5; i++ {
		if hint[i] == 'y' { // yellow
			possibles[i][AsciiToIndex(guess[i])] = false
			seen[AsciiToIndex(guess[i])] = true
		}
	}
	for i := 0; i < 5; i++ {
		if hint[i] != 'g' && hint[i] != 'y' { // we treat everything else as grey
			for j := 0; j < 5; j++ {
				if j == i || !seen[AsciiToIndex(guess[i])] {
					possibles[j][AsciiToIndex(guess[i])] = false
				}
			}
		}
	}
}
