package meetup

import (
	"fmt"
	"time"
)

type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second              = "second"
	Third               = "third"
	Fourth              = "fourth"
	Teenth              = "teenth"
	Last                = "last"
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	fmt.Printf("Want %v %v of %v\n", week, weekday, month)

	for day := 1; day <= 31; day++ {
		candidate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		fmt.Printf("%v %d is a %v\n", month, candidate.Day(), candidate.Weekday())
		if candidate.Weekday() == weekday && isMatchingWeekSchedule(week, candidate) {
			return candidate.Day()
		}
	}
	panic("no solution found")
}

func isMatchingWeekSchedule(week WeekSchedule, candidate time.Time) bool {
	switch week {
	case First:
		return candidate.Day() <= 7
	case Second:
		return candidate.Day() >= 8 && candidate.Day() <= 14
	case Third:
		return candidate.Day() >= 15 && candidate.Day() <= 21
	case Fourth:
		return candidate.Day() >= 22 && candidate.Day() <= 28
	case Teenth:
		return candidate.Day() > 12 && candidate.Day() < 20
	case Last:
		lastDay := time.Date(candidate.Year(), candidate.Month()+1, 0, 0, 0, 0, 0, time.UTC)
		fmt.Printf("lastDay of month is %v\n", lastDay.Day())
		return candidate.Day() > lastDay.Day()-7
	}
	return false
}
