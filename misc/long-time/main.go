package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	timeDividerByLevel = map[int]float64{
		1: 60,
		2: 60,
		3: 24,
		4: 365,
	}

	longUnitByLevel = map[int]string{
		1: "second",
		2: "minute",
		3: "hour",
		4: "day",
		5: "year",
	}
)

func secondsToLongTime(seconds int) string {
	level := 1
	var ok bool
	var divider float64
	timeU := float64(seconds)

	for {
		divider, ok = timeDividerByLevel[level]
		if !ok {
			//todo: check what should be returned here
			return ""
		}

		//At this point, "timeU" will have reached it's largest time unit.
		if divider > timeU {
			break
		}

		timeU = timeU / divider
		level++
	}

	var longU string
	timeUnits := make([]string, 0, level)
	//	1.36
	for {
		longU, ok = longUnitByLevel[level]
		if !ok {
			//todo: check what should be returned here
			return ""
		}

		if math.Floor(timeU) > 1 {
			longU += "s"
		}

		timeUnits = append(timeUnits, fmt.Sprintf("%d %s", int(timeU), longU))

		level--

		//No more decimal part left
		if math.Floor(timeU) == timeU || level == 0 {
			break
		}

		divider, ok = timeDividerByLevel[level]
		if !ok {
			//todo: check what should be returned here
			return ""
		}
		timeU = (timeU * divider) - (math.Floor(timeU) * divider)
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

func roundTwo(x float64) float64 {
	return math.Round(x*100) / 100
}

func main() {
	var x = 5.0
	var y = 5.0
	fmt.Println(math.Floor(x) == y) // outputs "5"
}
