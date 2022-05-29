package main

import "fmt"

func showhelp() {
	fmt.Println(
		"wordleg [--help] [--showlogic] guess1 result1 guess2 result2 ...\n"+
			"\n"+
			"This program suggests guesses for Wordle (the NYT variant) based on the number\n"+
			"  of hint groups the guess would split the remaining possibilities into.\n"+
			"\n"+
			"Guesses and results must be in pairs, and all in lowercase.\n"+
			"\n"+
			"Options:\n"+
			"\n"+
			"  --help\n"+
			"\n"+
			"    Shows this message.\n"+
			"\n"+
			"  --showlogic\n"+
			"\n"+
			"    Shows all of the filtering steps taken, and up to 100 of the remaining\n"+
			"      valid solutions.\n",
		"",
	)
}

func showversion() {
	fmt.Println(
		"wordleg b.24 - 29 May 22 - github.com/chriszuelsdorf",
	)
}
