package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("aspirations.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var aspirations []string
	if err := json.Unmarshal(bytes, &aspirations); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	total := len(aspirations)
	if total == 0 {
		fmt.Println("No aspirations found!")
		return
	}

	now := time.Now()
	dayOfYear := now.YearDay()
	hour := now.Hour()

	// Generate two indices based on both day of year and hour
	index1 := (dayOfYear*24 + hour) % total
	index2 := (hour*dayOfYear + 17) % total // 17 is just a random prime to offset

	fmt.Printf("Aspirations for Day %d, Hour %02d:\n\n", dayOfYear, hour)
	fmt.Printf("%d. %s\n", index1+1, aspirations[index1])
	// Avoid duplicate if both indices happen to be the same
	if index2 != index1 {
		fmt.Printf("%d. %s\n", index2+1, aspirations[index2])
	}
}