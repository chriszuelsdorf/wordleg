package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func genall(current_length int) []string {
	l := "abcdefghijklmnopqrstuvwxyz"
	o := make([]string, 0, int(math.Pow(26, float64(current_length))))
	if current_length == 1 {
		for _, v := range l {
			o = append(o, string(v))
		}
	} else {
		for _, v := range genall(current_length - 1) {
			for _, letter := range l {
				o = append(o, v+string(letter))
			}
		}
	}
	return o
}

func do_main() ([]string, error) {
	// commenting this out bc it's a valid use case
	// if len(os.Args) == 1 {
	// 	return []string{}, errors.New("You haven't supplied any words!")
	// }
	ags := os.Args[1:]
	showall := false
	k := 0
	for _, v := range ags {
		if strings.HasPrefix(v, "--") {
			if v == "--showlogic" {
				showall = true
			} else if v == "--help" {
				showhelp()
				return []string{}, errors.New("")
			} else if v == "--version" {
				showversion()
				return []string{}, errors.New("")
			} else {
				return []string{}, errors.New("Unknown switch " + v + ".")
			}
		} else {
			if len(v) != 5 {
				return []string{}, errors.New("Argument " + v + " was not 5 characters long")
			}
			if k%2 == 1 {
				for _, l := range v {
					if !strings.ContainsRune("byg", l) {
						return []string{}, errors.New("Argument " + v + " was interpreted as a match-pattern but contained a non-match letter!")
					}
				}
			}
			ags[k] = v
			k++
		}
	}
	ags = ags[:k]
	possibs := getsolns()
	for i := 0; i < len(ags)/2; i++ {
		guess := ags[2*i]
		result := ags[(2*i)+1]
		if showall {
			fmt.Println("Eval guess " + strconv.Itoa(i*2) + " - guess was " + string(guess) + " and result was " + string(result))
		}
		if len(guess) != 5 || len(result) != 5 {
			return []string{}, errors.New("Make sure that your guesses and results are both 5-long! Issue was with `" + guess + "` and `" + result + "`!")
		}
		for j := 0; j < 5; j++ {
			if showall {
				fmt.Println("--> Eval letter " + strconv.Itoa(j) + " - guess was " + string(guess[j]) + " and result was " + string(result[j]))
			}
			if result[j] == 'g' {
				k := 0
				for _, v := range possibs {
					if v[j] == guess[j] {
						possibs[k] = v
						k++
					}
				}
				possibs = possibs[:k]
				if showall {
					fmt.Println("    --> Result was g; " + strconv.Itoa(k) + " matched.")
				}
			} else if result[j] == 'b' {
				if strings.IndexByte(guess, guess[j]) == j {
					k := 0
					for _, v := range possibs {
						if strings.IndexByte(v, guess[j]) == -1 {
							possibs[k] = v
							k++
						}
					}
					possibs = possibs[:k]
					if showall {
						fmt.Println("    --> Result was b and it was the first; " + strconv.Itoa(k) + " matched.")
					}
				} else {
					k := 0
					for _, v := range possibs {
						if strings.Count(v, string(guess[j])) < strings.Count(guess[:j+1], string(guess[j])) {
							possibs[k] = v
							k++
						}
					}
					possibs = possibs[:k]
					if showall {
						fmt.Println("    --> Result was b and it wasn't the first; " + strconv.Itoa(k) + " matched.")
					}
				}
			} else if result[j] == 'y' {
				k := 0
				for _, v := range possibs {
					if strings.IndexByte(v, guess[j]) != -1 && strings.IndexByte(v, guess[j]) != j {
						possibs[k] = v
						k++
					}
				}
				possibs = possibs[:k]
				if showall {
					fmt.Println("    --> Result was y; " + strconv.Itoa(k) + " matched.")
				}
			}
		}
	}
	if showall {
		fmt.Println("")
	}
	fmt.Print(strconv.Itoa(len(possibs)) + " remaining")
	if showall {
		fmt.Println(": ")
		fmt.Println(possibs[:int(math.Min(float64(len(possibs)), 100))])
	} else {
		fmt.Println(".")
	}
	fmt.Println("")
	return possibs, nil
}

func main() {
	fmt.Println("")
	t, err := do_main()
	if err == nil {
		fmt.Println("Showing recommendations:")
		t2 := diagit(t)
		for i, v := range t2 {
			if i < 25 {
				fmt.Println("--> " + v.word + " - groups: " + strconv.Itoa(v.groups))
			}
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println("")
}
