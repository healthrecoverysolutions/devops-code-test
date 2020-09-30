package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"fmt"
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

		oldVal, olderr := strconv.Atoi(record[1])
		if olderr != nil {
			log.Fatal(olderr)
		}

		newVal, newerr := strconv.Atoi(record[2])
		if newerr != nil {
			log.Fatal(newerr)
		}

		if newerr == nil && olderr == nil && (newVal > oldVal) {
			fmt.Printf("%s %d\n", record[0], newVal - oldVal)
		}
	}
}
