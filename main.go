package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	currentTime := time.Now()

	startTime, err := parseTimeString("08:00:00", currentTime)
	if err != nil {
		panic("Failed to parse time")
	}

	endTime, err := parseTimeString("13:00:00", currentTime)
	if err != nil {
		panic("Failed to parse time")
	}

	if len(os.Args) >= 2 {

		switch os.Args[1] {
		case "futureDiff":
			diff := endTime.Sub(currentTime)

			fmt.Print(diff.Hours())

		case "pastDiff":
			fmt.Println(currentTime.Sub(startTime))

		}

	} else {
		fmt.Println("Supply arguments: Possible arguments:")
		fmt.Println("futureDiff")
		fmt.Println("pastDiff")
	}

}

func parseTimeString(inputString string, currentTime time.Time) (time.Time, error) {
	time, err := time.ParseInLocation(time.TimeOnly, inputString, currentTime.Location())
	// Set the date as the current date
	time = time.AddDate(currentTime.Year(), int(currentTime.Month())-1, currentTime.Day()-1)

	return time, err
}
