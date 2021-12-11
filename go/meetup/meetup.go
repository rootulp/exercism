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
	fmt.Printf("Want weekday %v and week %v\n", weekday, week)

	for day := 0; day <= 31; day++ {
		candidate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		fmt.Printf("candidate.Weekday()=%v\n", candidate.Weekday())
		if candidate.Weekday() == weekday && isMatchingWeekSchedule(week, candidate) {
			return candidate.Day()
		}
	}
	panic("no solution found")
}

func isMatchingWeekSchedule(week WeekSchedule, candidate time.Time) bool {
	switch week {
	case Last:
		return isLast(candidate)
	case Teenth:
		return isTeenth(candidate)
	case First:
		return candidate.Day() <= 7
	case Second:
		return candidate.Day() >= 8 && candidate.Day() <= 14
	case Third:
		return candidate.Day() >= 15 && candidate.Day() <= 21
	case Fourth:
		return candidate.Day() >= 22 && candidate.Day() <= 28
	}
	return false
}

func isTeenth(candidate time.Time) bool {
	return candidate.Day() > 12 && candidate.Day() < 20
}

func isLast(candidate time.Time) bool {
	// lastDay := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	// if t.Day() == lastDay {
	// 	return Last
	// }
	return false
}
