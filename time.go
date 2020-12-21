package rdg

import (
	"math/rand"
	"time"
)

// Time generates a random time
func Time() time.Time {
	var t time.Time
	return t.Add(time.Duration(rand.Int63()))
}

// TimeBetween generates a random time between two times.
// if the start time is after the end time, the generated
// time will be before the start time.
func TimeBetween(start, end time.Time) time.Time {
	if start.Equal(end) {
		return start
	}
	sy, sm, sd := start.Date()
	sh := start.Hour()
	ss := start.Second()
	sn := start.Nanosecond()

	ey, em, ed := end.Date()
	eh := end.Hour()
	es := end.Second()
	en := end.Nanosecond()

	year := IntBetween(sy, ey)
	month := IntBetween(int(sm), int(em))
	day := IntBetween(sd, ed)
	hour := IntBetween(sh, eh)
	second := IntBetween(ss, es)
	nano := IntBetween(sn, en)

	t := start.AddDate(year-1, month-1, day-1)
	t = t.Add(time.Hour * time.Duration(hour))
	t = t.Add(time.Second * time.Duration(second))
	t = t.Add(time.Duration(nano))
	return t
}

// TimeWithin generates a random time with a start time a specified duration.
// if the duration is negative the generated time will be before the start.
func TimeWithin(start time.Time, end time.Duration) time.Time {
	return TimeBetween(start, start.Add(end))
}
