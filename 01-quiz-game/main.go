package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	// Parse flags
	problems := flag.String("f", "problems.csv", "Path to problems csv file")
	limit := flag.Int("limit", 30, "Seconds until the quiz times out")
	flag.Parse()

	// Read csv
	f, err := os.Open(*problems)
	if err != nil {
		panic(fmt.Sprintf("File: %s not found", *problems))
	}

	// Game and timeout threads
	r := csv.NewReader(f)
	n := 0
	points := 0

	stop := make(chan bool, 1)

	go func() {
		time.Sleep(time.Duration(*limit) * time.Second)

		// Run through remaining
		for {
			_, err := r.Read()
			if err == io.EOF {
				break
			}
			n++
		}
		fmt.Println()
		fmt.Println("You timed out!")
		stop <- true
	}()

	go func() {
		var userAnswer string
		for {
			record, err := r.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			n++
			fmt.Printf("%d) %s: ", n, record[0])
			fmt.Scanf("%s\n", &userAnswer)

			if userAnswer == record[1] {
				points++
			}
		}
		stop <- true
	}()

	<-stop
	fmt.Printf("Scored %d/%d point(s)", points, n)
}
