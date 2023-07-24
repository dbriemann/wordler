package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"briemann.com/wordler/algo"
	"briemann.com/wordler/dict"
	"briemann.com/wordler/rules"
)

func main() {
	hardmode := flag.Bool("hardmode", false, "enable this to use only reduced set of words")
	secret := flag.String("secret", "", "set a secret to be solved")
	verbose := flag.Bool("verbose", false, "enables verbose logging of info")

	flag.Parse()

	possibles := algo.WordSet{}
	for i := 0; i < 5; i++ {
		for c := 'a'; c <= 'z'; c++ {
			possibles[i][algo.AsciiToIndex(byte(c))] = true
		}
	}
	remaining := make([]string, len(dict.Words))
	copy(remaining, dict.Words)

	if *hardmode {
		// Overwrite with smaller set.
		algo.ValidGuesses = dict.Words
	}

	if *secret != "" {
		valid := false
		if len(*secret) == 5 {
			for _, w := range dict.Words {
				if w == *secret {
					valid = true
					break
				}
			}
		}
		if !valid {
			panic("invalid secret provided")
		}

		start := time.Now()
		// Run complete game.
		for i := 1; ; i++ {
			guess := algo.Guess(&possibles, &remaining, i, false)
			fmt.Printf("Guess %d is: %q\n", i, guess)
			hint := rules.Check(guess, *secret)
			if hint == "ggggg" {
				fmt.Printf("won on attempt %d\n", i)
				break
			}
			algo.Reduce(&possibles, guess, hint)
		}
		fmt.Println("Time used:", time.Now().Sub(start))
		return
	}

	// Else play an interactive game.
	reader := bufio.NewReader(os.Stdin)
	for i := 1; ; i++ {
		guess := algo.Guess(&possibles, &remaining, i, *verbose)
		fmt.Printf("Guess %d is: %q\n", i, guess)
		fmt.Printf("Enter hint in short notation (g(reen) / y(ellow) / x(gray)), e.g. ggxyx: \n")
		hint, err := reader.ReadString('\n')
		if err != nil {
			panic("error")
		}
		hint = strings.TrimSpace(hint)
		if hint == "ggggg" {
			fmt.Printf("won on attempt %d\n", i)
			return
		}
		algo.Reduce(&possibles, guess, hint)
	}
}
