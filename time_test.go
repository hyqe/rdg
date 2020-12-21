package rdg

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	x := Time()
	if x.IsZero() {
		t.Fatal("generated zero time")
	}
}

func TestTimeBetween(t *testing.T) {
	t.Parallel()
	t.Run("one_option", func(t *testing.T) {
		now := time.Now()
		x := TimeBetween(now, now)
		if x != now {
			t.Fatalf("want: %v got: %v", now, x)
		}
	})
	t.Run("start_after_end", func(t *testing.T) {
		start := time.Time{}.Add(time.Hour * 2)
		end := time.Time{}.Add(time.Hour * 1)
		x := TimeBetween(start, end)
		if start.Before(x) && start.After(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("low", func(t *testing.T) {
		start := time.Time{}
		end := time.Time{}.Add(1)
		x := TimeBetween(start, end)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("mid", func(t *testing.T) {
		start := time.Time{}
		end := time.Time{}.Add(92233720)
		x := TimeBetween(start, end)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("high", func(t *testing.T) {
		start := time.Time{}
		end := time.Time{}.Add(9223372036854775805)
		x := TimeBetween(start, end)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
}

func TestTimeWthin(t *testing.T) {
	t.Parallel()
	t.Run("one_option", func(t *testing.T) {
		now := time.Now()
		x := TimeWithin(now, 0)
		if x != now {
			t.Fatalf("want: %v got: %v", now, x)
		}
	})
	t.Run("start_after_end", func(t *testing.T) {
		duration := time.Hour
		start := time.Time{}.Add(time.Hour * 2)
		x := TimeWithin(start, time.Hour)
		if start.Before(x) && start.After(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, start.Add(duration), x)
		}
	})
	t.Run("low", func(t *testing.T) {
		duration := time.Duration(1)
		start := time.Time{}
		end := time.Time{}.Add(duration)
		x := TimeWithin(start, duration)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("mid", func(t *testing.T) {
		duration := time.Duration(92233720)
		start := time.Time{}
		end := time.Time{}.Add(duration)
		x := TimeWithin(start, duration)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("high", func(t *testing.T) {
		duration := time.Duration(9223372036854775805)
		start := time.Time{}
		end := time.Time{}.Add(duration)
		x := TimeWithin(start, duration)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
}
