package rdg

import (
	"regexp"
	"strings"
	"testing"
)

func TestStreet(t *testing.T) {
	t.Run("not_blank", func(t *testing.T) {
		street := Street()
		if strings.TrimSpace(street) == "" {
			t.Fatalf("failed to generate value")
		}
	})
}

func TestStreet1(t *testing.T) {
	t.Run("not_blank", func(t *testing.T) {
		street := Street1()
		if strings.TrimSpace(street) == "" {
			t.Fatalf("failed to generate value")
		}
	})
	t.Run("match_regex", func(t *testing.T) {
		streetRegStr := `^\d{1,5}(:? \w+){0,4}$`
		streetReg := regexp.MustCompile(streetRegStr)
		street := Street1()
		if !streetReg.Match([]byte(street)) {
			t.Fatalf("want: %v got: %v ", streetRegStr, street)
		}
	})
}

func TestStreet2(t *testing.T) {
	t.Run("not_blank", func(t *testing.T) {
		street := Street2()
		if strings.TrimSpace(street) == "" {
			t.Fatalf("failed to generate value")
		}
	})
	t.Run("match_regex", func(t *testing.T) {
		streetRegStr := `(:?\w|[#])+ (:?[A-Z]|\d{1,2})$`
		streetReg := regexp.MustCompile(streetRegStr)
		street := Street2()
		if !streetReg.Match([]byte(street)) {
			t.Fatalf("want: %v got: %v ", streetRegStr, street)
		}
	})
}

func TestCity(t *testing.T) {
	t.Run("not_blank", func(t *testing.T) {
		x := City()
		if strings.TrimSpace(x) == "" {
			t.Fatalf("failed to generate value")
		}
	})
	t.Run("lookup", func(t *testing.T) {
		x := City()
		if !strArrContains(cities, x) {
			t.Fatalf("failed to find value: %v", x)
		}
	})
}

func TestState(t *testing.T) {
	t.Parallel()
	t.Run("not_blank", func(t *testing.T) {
		x := State()
		if strings.TrimSpace(x) == "" {
			t.Fatalf("failed to generate value")
		}
	})
	t.Run("lookup", func(t *testing.T) {
		x := State()
		if !strArrContains(states, x) {
			t.Fatalf("failed to find value: %v", x)
		}
	})
}

func TestZip(t *testing.T) {
	t.Parallel()
	t.Run("not_blank", func(t *testing.T) {
		x := Zip()
		if strings.TrimSpace(x) == "" {
			t.Fatalf("failed to generate value")
		}
	})
}

func TestCounrty(t *testing.T) {
	t.Parallel()
	t.Run("not_blank", func(t *testing.T) {
		x := Country()
		if strings.TrimSpace(x) == "" {
			t.Fatalf("failed to generate value")
		}
	})
	t.Run("lookup", func(t *testing.T) {
		x := Country()
		if !strArrContains(countries, x) {
			t.Fatalf("failed to find value: %v", x)
		}
	})
}

func strArrContains(a []string, x string) bool {
	found := false
	for _, s := range a {
		if s == x {
			found = true
			break
		}
	}
	if found {
		return true
	}
	return false
}
