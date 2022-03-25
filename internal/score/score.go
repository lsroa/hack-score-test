package score

import (
	"encoding/csv"
	"os"
)

type Row struct {
	Name  string
	Base  string
	Input string
}

func GenerateTable(pathFile string) []Row {

	file, err := os.Open(pathFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	table, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	output := make([]Row, len(table))

	for index, row := range table {
		output[index] = Row{Name: row[0], Base: row[1], Input: row[2]}
	}

	return output
}
