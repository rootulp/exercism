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
	time := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Want weekday %v and week %v so weekSchedule(%v) %v\n", weekday, week, time, weekSchedule(time))

	for time.Weekday() != weekday || weekSchedule(time) != week {
		fmt.Printf("time.Weekday()=%v and weekSchedule(%v)=%v\n", time.Weekday(), time, weekSchedule(time))
		time = time.AddDate(0, 0, 1) // move forward one day
	}
	fmt.Printf("%v\n", time)
	return time.Day()
}

// TODO this weekSchedule is incorrect.
// This should be based on the # of mondays in a month.
func weekSchedule(t time.Time) WeekSchedule {
	if t.Day() == 1 {
		return First
	}
	if t.Day() == 2 {
		return Second
	}
	if t.Day() == 3 {
		return Third
	}
	if t.Day() == 4 {
		return Fourth
	}
	if t.Day() > 12 && t.Day() < 20 {
		return Teenth
	}
	lastDay := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	if t.Day() == lastDay {
		return Last
	}
	return "none"
}
