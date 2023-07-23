# wordler

solves wordle

# main

To build the solver just run `go build` in the root directory.

By default `wordler` will start an interactive mode (input via stdin) and will use all allowed words (by Wordle) for guessing.

You can restrict the words by providing the `-hardmode` flag, to only allow valid solutions as guesses. This will make it much faster but it will need more guesses for some words.

You can specify a secret which will then be instantly solved in a non-interactive way by providing the `-secret <word>` flag.

Example output:

```bash
➜ ./wordler -hardmode -secret crust
Guess 1 is: "tares"
Guess 2 is: "noisy"
Guess 3 is: "blurt"
Guess 4 is: "crust"
won on attempt 4
Time used: 2.24170733s
```

# bench

To build the benchmark tool just run `go build` in the `bench` directory.

With the `-runs` flag you can specify how many words will be solved.

Same as for main you can enable hard mode via the `-hardmode` flag.
