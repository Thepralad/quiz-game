package main

import (
	"encoding/csv"
	"fmt"
	"flag"
	"os"
	"time"
)

func main(){
	//defining a flag for the CSV file name
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format fo 'question,answert'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()
		
	data := readCsv(*csvFileName)

	problem := parseData(data)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problem{
		fmt.Printf("Problem #%v: %s = ",i, p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
			case <-timer.C:
				fmt.Printf("\nScore: %v/%v\n", correct, len(problem))
				return
			case answer := <-answerCh: 
				if answer == p.a{
					correct++
				}
			}
	}
}

func parseData(line [][]string) []problem{
	ret := make ([]problem, len(line))
	for i, p := range line{
		ret[i] = problem{
			q: p[0],
			a: p[1],
		}
	}
	return ret
}

func readCsv(csvFileName string) [][]string{	
	//opening the CSV file
	fileName, err := os.Open(csvFileName)
	if err != nil{
		fmt.Printf("Can't open %v\n", csvFileName)
		os.Exit(1)
	}
	
	//reading the CSV file
	reader := csv.NewReader(fileName)
	data, err := reader.ReadAll()
	if err != nil{
		fmt.Printf("Can't open %v\n", csvFileName)
		os.Exit(1)
	}	
	return data
}

type problem struct{
	q string
	a string
}
