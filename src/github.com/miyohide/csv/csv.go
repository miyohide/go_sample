package main

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"encoding/csv"
	"io"
	"flag"
	"log"
	"os"
	"fmt"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	flag.Parse()

	file1, err := os.Open(flag.Arg(0))
	failOnError(err)
	defer file1.Close()

	file2, err := os.Create(flag.Arg(1))
	failOnError(err)
	defer file2.Close()

	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	writer := csv.NewWriter(transform.NewWriter(file2, japanese.ShiftJIS.NewEncoder()))
	writer.UseCRLF = true

	log.Printf("Start")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}
		var new_record []string
		for i, v := range record {
			if i >= 0 {
				new_record = append(new_record, fmt.Sprint(i) + ":" + v)
			}
		}
		writer.Write(new_record)
	}
	writer.Flush()
	log.Printf("Finish !")
}
