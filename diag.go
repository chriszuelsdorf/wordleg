package main

import (
	// "fmt"
	"sort"

	// "strconv"
	"strings"
)

func dstring(real_word string, guess_word string) string {
	out_string := ""
	for i, v := range guess_word {
		if real_word[i] == guess_word[i] {
			out_string += "g"
		} else if strings.Count(real_word, string(v)) > strings.Count(guess_word[:i], string(v)) {
			out_string += "y"
		} else {
			out_string += "b"
		}
	}
	return out_string
}

type Possib struct {
	word   string
	groups int
}

func diagit(possibs []string) []Possib {
	out := []Possib{}
	// mxgrps := 0
	for _, gw := range possibs {
		grps := ""
		for _, rw := range possibs {
			t := dstring(rw, gw)
			if !strings.Contains(grps, t) {
				grps += "," + t + ","
			}
		}
		// ngrps := strings.Count(grps, ",") / 2
		// fmt.Println(gw + " - " + strconv.Itoa(ngrps) + " - " + grps)
		out = append(out, Possib{gw, strings.Count(grps, ",") / 2})
		// if ngrps > mxgrps {
		// 	fmt.Println("New max groups!")
		// 	mxgrps = ngrps
		// }
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].groups > out[j].groups {
			return true
		} else if out[i].groups == out[j].groups {
			return out[i].word > out[j].word
		} else {
			return false
		}
	})
	return out
}
