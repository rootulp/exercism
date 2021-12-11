package meetup

import (
	"fmt"
	"time"
)

type WeekSchedule string

const First WeekSchedule = "first"
const Second WeekSchedule = "second"
const Third WeekSchedule = "third"
const Fourth WeekSchedule = "fourth"
const Teenth WeekSchedule = "teenth"
const Last WeekSchedule = "last"

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	for day := 1; day <= 31; day++ {
		candidate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		// fmt.Printf("%v %d is a %v\n", month, candidate.Day(), candidate.Weekday())
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
		lastDayOfMonth := time.Date(candidate.Year(), candidate.Month()+1, 0, 0, 0, 0, 0, time.UTC)
		return candidate.Day() > lastDayOfMonth.Day()-7
	}
	panic(fmt.Sprintf("%v did not match a WeekSchedule", week))
}
