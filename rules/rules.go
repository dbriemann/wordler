package rules

import (
	"math"
)

// Hints from Wordle are basically ternary numbers.
// Each place can be either:
// 0 (grey): character is not contained in the word,
// 1 (yellow): character is contained but in the wrong spot,
// 2 (green): character is correctly place.
// Since every word has 5 characters a hint can be encoded in a byte.
type Hint uint8

func (h Hint) ToColors() string {
	colors := [5]byte{}

	for i := 0; i < 5; i++ {
		rem := h % 3
		h = h / 3
		switch rem {
		case 0:
			colors[i] = 'x'
		case 1:
			colors[i] = 'y'
		case 2:
			colors[i] = 'g'
		}
	}

	return string(colors[:])
}

func HintFromColors(colors string) Hint {
	var h Hint

	for i := 0; i < 5; i++ {
		color := colors[i]
		switch color {
		case 'g':
			h += Hint(math.Pow(3, float64(i))) * 2
		case 'y':
			h += Hint(math.Pow(3, float64(i))) * 1
		default:
			// nothing to add
		}
	}

	return h
}

// Check implements the Wordle hinting. It returns a color code
// for a given secret and check. The color can be
// green(g): a character is correctly placed,
// yellow(y): a character is contained but placed elsewhere,
// grey(x): a character is not contained in the word.
//
// e.g.
//
// secret: aorta
//
//	guess: adult
//	 code: gxxxy
func Check(guess, secret string) string {
	hints := []byte{'x', 'x', 'x', 'x', 'x'}
	secretHandled := make([]bool, 5)
	guessHandled := make([]bool, 5)

	// Find all correctly placed characters (green).
	for i := 0; i < 5; i++ {
		if !secretHandled[i] && guess[i] == secret[i] {
			hints[i] = 'g'
			secretHandled[i] = true
			guessHandled[i] = true
		}
	}

	// Find all characters that are in the wrong place (yellow).
	for i := 0; i < 5; i++ {
		if !guessHandled[i] {
			for j := 0; j < 5; j++ {
				if !secretHandled[j] && i != j && guess[i] == secret[j] {
					hints[i] = 'y'
					guessHandled[i] = true
					secretHandled[j] = true
					break
				}
			}

		}
	}

	// Rest is already marked as non-existant (grey).

	return string(hints)
}
