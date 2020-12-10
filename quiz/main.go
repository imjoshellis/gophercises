package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("file", "problems.csv", "Filename for the quiz")
	durPtr := flag.Int("timer", 30, "Timer duration in seconds (default: 30)")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalf("Failed to open %v\n", *csvFilename)
	}
	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse the file. Are you sure it was csv?\n")
	}

	type problem struct {
		q string
		a string
	}

	parseLines := func(lines [][]string) []problem {
		res := make([]problem, len(lines))
		for i, line := range lines {
			res[i] = problem{
				q: line[0],
				a: strings.TrimSpace(line[1]),
			}
		}
		return res
	}
	problems := parseLines(lines)

	count := 0

	duration := time.Duration(*durPtr) * time.Second
	timer := time.NewTimer(duration)
	defer timer.Stop()

	for i := 0; i < len(problems); {
		fmt.Print(problems[i].q, "= ")
		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTimes up!")
			fmt.Printf("You got %v/%v problems.\n", count, len(problems))
			return
		case ans := <-ansCh:
			if strings.TrimSpace(ans) == problems[i].a {
				fmt.Println("Correct!")
				count++
				i++
			} else {
				fmt.Println("Incorrect, try again...")
			}
		}
	}
}
