package strats

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Load strategy csv, hard total, soft total, split
func Load() (map[int]map[int]string, map[int]map[int]string, map[int]map[int]string) {

	return csvToMap(loadCsv("strats/hard_total.csv")), csvToMap(loadCsv("strats/soft_total.csv")), csvToMap(loadCsv("strats/splitting.csv"))

}

func loadCsv(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	csvFile, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	return csvFile
}

func csvToMap(file [][]string) map[int]map[int]string {

	stratMap := make(map[int]map[int]string)

	colNames := file[0][1:]
	for _, handValue := range file[1:] {

		playerTotal, err := strconv.Atoi(handValue[0])
		if err != nil {
			panic(err)
		}

		stratMap[playerTotal] = make(map[int]string)

		for i, move := range handValue[1:] {

			dealerUpCard, err := strconv.Atoi(colNames[i])
			if err != nil {
				panic(err)
			}
			stratMap[playerTotal][dealerUpCard] = move
		}
	}

	return stratMap
}
