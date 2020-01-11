package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Parse flags
	problems := flag.String("f", "problems.csv", "Path to problems csv file")
	flag.Parse()

	// Read csv
	f, err := os.Open(*problems)
	if err != nil {
		panic(fmt.Sprintf("File: %s not found", *problems))
	}

	// Game loop
	points := 0
	var userAnswer string
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s: ", record[0])
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == record[1] {
			points++
		}
	}

	fmt.Printf("You got %d point(s)!", points)
}
