package main

import (
	"fmt"
	"strings"
)

var (
	modulus   = []int{60, 60, 24, 365}
	longUnits = []string{"second", "minute", "hour", "day", "year"}
)

func secondsToLongTime(seconds int) string {
	var remainder int
	timeU := seconds
	timeParts := make([]int, 0)

	for _, modOp := range modulus {
		//At this point, "timeU" will have reached it's largest time unit.
		if modOp > timeU {
			break
		}

		remainder = timeU % modOp
		timeParts = append(timeParts, remainder)
		timeU = (timeU - remainder) / modOp
	}

	timeParts = append(timeParts, timeU)

	var longU string
	timeUnits := make([]string, 0, len(timeParts))

	for i := len(timeParts) - 1; i >= 0; i-- {
		longU = longUnits[i]
		timeU = timeParts[i]

		if timeU > 1 {
			longU += "s"
		}

		if timeU == 0 && i == 0 {
			continue
		}

		timeUnits = append(timeUnits, fmt.Sprintf("%d %s", timeU, longU))
	}

	return convertToString(timeUnits)
}

func convertToString(timeUnits []string) string {
	if len(timeUnits) == 1 {
		return timeUnits[0]
	}

	s := strings.Join(timeUnits[:len(timeUnits)-1], ", ")

	return fmt.Sprintf("%s and %s", s, timeUnits[len(timeUnits)-1])
}
