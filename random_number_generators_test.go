package rdg

import (
	"testing"
	"time"
	"math/rand"
)

func TestRandomNumberGenerators(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var maxInt32 int32 = 2147483647
	var maxInt64 int64 = 9223372036854775807

	for i := 1; i < 1000; i++ {
		maxInt32 = maxInt32 - int32(float32(maxInt32)*.0175)
		maxInt64 = maxInt64 - int64(float64(maxInt64)*.04)
		//fmt.Printf("int32=(%v)(%v), int64=(%v)(%v)\n", maxInt32, maxInt32 * -1, maxInt64, maxInt64 * -1)


		int32value := NewInt32Between(maxInt32 * -1, maxInt32)
		int64value := NewInt64Between(maxInt64 * -1, maxInt64)
		if int32value < maxInt32 * -1 {
			t.Fatalf("Value returned went below Minimum Value!")
		}
		if int32value  > maxInt32 {
			t.Fatalf("Value returned went above Maximum Value!")
		}
		if int64value < maxInt64 * -1 {
			t.Fatalf("Value returned went below Minimum Value!")
		}
		if int64value > maxInt64 {
			t.Fatalf("value returned went above Maximum Value!")
		}


	}

}
