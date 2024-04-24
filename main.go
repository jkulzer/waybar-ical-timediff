package main

import (
	"fmt"
	"github.com/arran4/golang-ical"
	"math"
	"os"
	"strings"
	"time"
)

type TimeDuration time.Duration

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

			fmt.Println(TimeDuration(diff).Format("15:04"))

		case "pastDiff":
			diff := currentTime.Sub(startTime)

			fmt.Println(TimeDuration(diff).Format("15:04"))

		case "diffPercentage":
			startFromIcal()
			fmt.Println(fmt.Sprint(diffPercentage(currentTime, startTime, endTime)) + "%")
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

// see https://stackoverflow.com/a/69101998
func (t TimeDuration) Format(format string) string {
	return time.Unix(0, 0).UTC().Add(time.Duration(t)).Format(format)
}

func diffPercentage(currentTime time.Time, startTime time.Time, endTime time.Time) int {

	entireDiff := endTime.Sub(startTime).Hours()
	pastDiff := currentTime.Sub(startTime).Hours()

	return int(math.Round(((pastDiff / entireDiff) * 100)))
}

func startFromIcal() {
	dat, err := os.ReadFile("./calendar.ics")
	if err != nil {
		panic("Failed reading calendar file")
	}

	cal, err := ics.ParseCalendar(strings.NewReader(string(dat)))

	events := cal.Events()

	for _, event := range events {
		end, err := event.GetEndAt()
		if err != nil {
			panic("Failed parsing events")
		}
		fmt.Println(end)
	}

}
