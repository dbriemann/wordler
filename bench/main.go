package main

import (
	"flag"
	"fmt"
	"time"

	"briemann.com/wordler/algo"
	"briemann.com/wordler/dict"
	"briemann.com/wordler/rules"
	"github.com/pkg/profile"
)

const (
	MAXTRIES = 6
)

func main() {
	hardmode := flag.Bool("hardmode", false, "enable this to use only reduced set of words")
	runs := flag.Int("runs", 10, "set the maximum number of words to be solved")

	flag.Parse()

	if *hardmode {
		// Overwrite with smaller set.
		algo.ValidGuesses = dict.Words
	}

	results := map[int]int{}
	failed := []string{}

	start := time.Now()
	prof := profile.Start(profile.ProfilePath("."))
	defer prof.Stop()

	for s := 0; s < *runs; s++ {
		secret := dict.Words[s]

		possibles := algo.WordSet{}
		for i := 0; i < 5; i++ {
			for c := 'a'; c <= 'z'; c++ {
				possibles[i][algo.AsciiToIndex(byte(c))] = true
			}
		}
		remaining := make([]string, len(dict.Words))
		copy(remaining, dict.Words)

		guess := ""
		i := 1

		for {
			if i > MAXTRIES {
				fmt.Printf("Secret %q NOT guessed after %d attempts.\n", secret, i-1)
				break
			}
			guess = algo.Guess(&possibles, &remaining, i, false)
			if guess == secret {
				results[i]++
				break
			}
			hint := rules.Check(guess, secret)
			algo.Reduce(&possibles, guess, hint)
			i++
		}

		if i > MAXTRIES {
			failed = append(failed, secret)
		} else {
			fmt.Printf("Secret %q guessed after %d attempts.\n", secret, i)
		}
	}
	timeUsed := time.Now().Sub(start)
	fmt.Println("time used:", timeUsed)

	totalGuesses := 0
	totalWords := 0
	for i := 1; i <= MAXTRIES; i++ {
		fmt.Printf("%d: %d\n", i, results[i])
		totalGuesses += results[i] * i
		totalWords += results[i]
	}
	fmt.Printf("Failed: %d\n", len(failed))
	fmt.Printf("Succeeded: %d\n", *runs-len(failed))
	fmt.Printf("Total guesses: %d\n", totalGuesses)
	fmt.Printf("Average guesses: %f\n", float64(totalGuesses)/float64(totalWords))
}
