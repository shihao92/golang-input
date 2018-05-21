package main

import (
	"fmt"
	"strings"
	"sort"
	"strconv"
)

var ValueStore = []string{}
var RankWords = []string{}

type RankAnotherWords struct {
	Word  string
	Count int
}

func processInput(param string) {
	// rankAnotherWords := []RankAnotherWords{}
	ValueStore := strings.Split(param, " ")
	rankAnotherWords := []RankAnotherWords{}

	fmt.Println( ValueStore )

	for _, value := range ValueStore {
		// if length of rank words == 0 else check through rank store
		if len(rankAnotherWords) == 0 {
			rankAnotherWords = append(rankAnotherWords, RankAnotherWords{
				Word: value,
				Count: 1,
			})
		} else {
			// check whether words exist within ranking
			var stringExist = false
			for item := 0; item < len(rankAnotherWords); item++ {
				attr := &rankAnotherWords[item]
				if attr.Word == value {
					stringExist = true
					break
				}
			}
			if stringExist != true {
				rankAnotherWords = append(rankAnotherWords, RankAnotherWords{
					Word: value,
					Count: 0,
				})
			}

			for item := 0; item < len(rankAnotherWords); item++ {
				attr := &rankAnotherWords[item]
				if attr.Word == value {
					attr.Count = attr.Count + 1
				}
			}

		}
	}
	// Transverse the rank struct and sort out the rank
	sort.Slice(rankAnotherWords, func(i, j int) bool { return rankAnotherWords[i].Count > rankAnotherWords[j].Count })
	// for display
	if( len( rankAnotherWords ) < 10 ) {
		for item := 0; item < len(rankAnotherWords); item++ {
			tmpCount := strconv.Itoa(item + 1)
			convertedCount := strconv.Itoa(rankAnotherWords[item].Count)
			fmt.Printf( "Rank number " + tmpCount + ": " );
			fmt.Printf( rankAnotherWords[item].Word );
			fmt.Println( " - " + convertedCount + " times" )
		}
	} else {
		for item := 0; item < 10; item++ {
			tmpCount := strconv.Itoa(item + 1)
			convertedCount := strconv.Itoa(rankAnotherWords[item].Count)
			fmt.Printf( "Rank number " + tmpCount + ": " );
			fmt.Printf( rankAnotherWords[item].Word );
			fmt.Println( " - " + convertedCount + " times" )
		}
	}
}

func main() {
	// place the text over here
	processInput("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
}