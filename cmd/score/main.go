package main

import (
	"fmt"
	"sort"

	"lsroa/hack-score/internal/math"
	"lsroa/hack-score/internal/score"
)

type Score struct {
	Player string
	Value  int
}

func main() {

	table := score.GenerateTable("score.csv")

	normalized_table := make([]Score, len(table))

	for index, row := range table {
		normalized_table[index] = Score{Player: row.Name, Value: math.FindValue(row.Base, row.Input)}
	}

	sort.Slice(normalized_table, func(i, j int) bool {
		return normalized_table[i].Value > normalized_table[j].Value
	})

	for _, row := range normalized_table {
		fmt.Println(row)
	}
}
