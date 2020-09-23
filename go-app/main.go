package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	if _, err := r.Read(); err != nil {
		panic(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		oldVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalln(err)
		}

		newVal, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalln(err)
		}

		if newVal > oldVal {
			fmt.Printf("%s %d\n", record[0], newVal - oldVal)
		}
	}
}
