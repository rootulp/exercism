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
	fmt.Printf("weekday %v and week %v so weekSchedule(%v) %v\n", weekday, week, time, weekSchedule(time))
	for time.Weekday() != weekday || weekSchedule(time) != week {
		fmt.Printf("weekSchedule(%v) %v\n", time, weekSchedule(time))
		time = time.AddDate(0, 0, 1) // move forward one day
	}
	fmt.Printf("%v\n", time)
	return time.Day()
}

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
	fmt.Printf("Last day in month %v is %v\n", t, lastDay)
	if t.Day() == lastDay {
		fmt.Printf("t.Day() = %v and lastDay %v", t.Day(), lastDay)
		return Last
	}
	return "none"
}
