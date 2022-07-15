package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	arr := []string{}
	file, err := os.Open("../example_data/poker.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	countP1Winner := PokerHand(arr)
	fmt.Println(countP1Winner)
}
