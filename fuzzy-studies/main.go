package main

import (
	"github.com/renstrom/fuzzysearch/fuzzy"
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
	"unicode/utf8"
)

func main()  {

	//fmt.Println(fuzzy.RankFind("LARRY TOMAS", []string{"LARRY THOMAS"}))
	fmt.Println(Levenshtein("LARRY THOMAS", "LARRY TOMAS"))

}


func Levenshtein(a, b string) int {
	f := make([]int, utf8.RuneCountInString(b)+1)

	for j := range f {
		f[j] = j
	}

	for _, ca := range a {
		j := 1
		fj1 := f[0] // fj1 is the value of f[j - 1] in last iteration
		f[0]++
		for _, cb := range b {
			mn := min(f[j]+1, f[j-1]+1) // delete & insert
			if cb != ca {
				mn = min(mn, fj1+1) // change
			} else {
				mn = min(mn, fj1) // matched
			}

			fj1, f[j] = f[j], mn // save f[j] to fj1(j is about to increase), update f[j] to mn
			j++
		}
	}

	return f[len(f)-1]
}
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func M1() {

	content, _ := ioutil.ReadFile("tn_customer_data_dev.txt")

	contentTxt := string(content[:])

	//fmt.Println(contentTxt)


	var names []string = strings.Split(contentTxt, "\n")

	fmt.Println(len(names))

	term := "NESTLE"

	matchedValues := fuzzy.Find(term, names)
	//
	//fmt.Println(matchedValues)

	rankingValues := fuzzy.RankFind(term, matchedValues)
	sort.Sort(rankingValues)

	for _,value:=range rankingValues {
		fmt.Println(value)
	}

}
