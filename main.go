package main

import (
	"errors"
	"fmt"
	"github.com/arran4/golang-ical"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

type TimeDuration time.Duration

func main() {
	currentTime := time.Now()

	if len(os.Args) >= 2 {

		switch os.Args[1] {
		case "futureDiff":
			_, endTime, err := diffsFromIcal(currentTime)
			if err != nil {
				os.Exit(0)
			}
			diff := endTime.Sub(currentTime)

			fmt.Println(TimeDuration(diff).Format("15:04"))

		case "diffPercentage":
			startTime, endTime, err := diffsFromIcal(currentTime)
			if err != nil {
				os.Exit(0)
			}
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

func getIcal() string {
	resp, err := http.Get("https://hwrical.zrgr.pw/informatik/semester2/kursa/Tutorium/Gruppe%201/Englisch/")
	if err != nil {
		panic("Failed reading calendar")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic("Failed to get calendar from webpage")
		}
		bodyString := string(bodyBytes)
		return string(bodyString)
	}

	return ""
}

func diffsFromIcal(currentTime time.Time) (time.Time, time.Time, error) {

	cal, err := ics.ParseCalendar(strings.NewReader(getIcal()))
	if err != nil {
		panic("Failed parsing calendar")
	}

	events := cal.Events()

	for _, event := range events {
		start, err := event.ComponentBase.GetStartAt()
		if err != nil {
			panic("Failed parsing events")
		}

		end, err := event.GetEndAt()
		if err != nil {
			panic("Failed parsing events")
		}

		if start.Before(currentTime) && end.After(currentTime) {
			return start, end, nil
		}

	}

	return time.Now(), time.Now(), errors.New("No current event found")

}
