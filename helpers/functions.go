package helpers

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func PickRandom(list *[][]string) string {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	randomIndex := random.Intn(len(*list))
	pick := (*list)[randomIndex]
	return pick[0]
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
