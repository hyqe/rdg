package rdg

import (
	"math/rand"
	"testing"
	"time"
)

var Tests = [][]interface{}{
	{100, 10.999, "Does this work?"},
	{uint(500), float64(9999.8888), "items"},
	{"more", "test", "data"},
}

func ifInterfaceIn(interfaceSlice []interface{}, object interface{}) bool {
	var boolean bool = false
	for _, value := range interfaceSlice {
		if object == value {
			boolean = true
		}
	}
	if !boolean {
		return false
	} else {
		return true
	}
}

func TestSelectAny(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		for _, test := range Tests {
			value, err := SelectAny(test...)
			if err != nil {
				t.Fatalf("no interfaces or slice/array of interfaces was passed\n")
			}

			if !ifInterfaceIn(test, value) {
				t.Fatalf("Failed to return a random interface from interfaces passed\n")
			}
		}
	}

}
