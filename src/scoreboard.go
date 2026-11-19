package src

import "fmt"

func ScoreBoard(line []Log) {
	count := map[string]int{}
	for _, l := range line {
		count[l.Hst]++
	}
	fmt.Println(count)
}
