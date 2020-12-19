package rdg

import (
	"math/rand"
	"time"
)

func NewInt32() int32 {
	rand.Seed(time.Now().UnixNano())
	value := rand.Int31n(2147483647)
	if rand.Intn(2) == 1 {
		value = value * -1
	}
	return value

}

func NewInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	value := rand.Int63n(9223372036854775807)
	if rand.Intn(2) == 1 {
		value = value * -1
	}
	return value
}

func NewInt32Between(min int32, max int32) int32 {

	// create an exact copy of min and max
	var minmax = map[string]int32{
		"min": min,
		"max": max,
	}
	// remove the negative sign if any exist
	for key, _ := range minmax {
		if minmax[key] < 0 {
			minmax[key] = minmax[key] * -1
		}
	}
	// create a variable to put the largest non negative value in from the "minmax map"
	var realmax int32
	if minmax["min"] > minmax["max"] {
		realmax = minmax["min"]
	} else {
		realmax = minmax["max"]
	}
	rand.Seed(time.Now().UnixNano())
	// use the realmax variable to attain a random value from rand.Int31n()
	value := rand.Int31n(realmax)
	// create a random chance that the return value is negative or positive
	if rand.Intn(2) == 1 {
		value = value * -1
	}
	// if we accidently create a value that is below the real minimum value(provided by the user)
	// redo the process until we get a value within the proper minimum and maximum bounds(provided by the user)
	for value < min || value > max {
		value = rand.Int31n(realmax)
		if rand.Intn(2) == 1 {
			value = value * -1
		}
	}
	return value
}

func NewInt64Between(min int64, max int64) int64 {

	// create an exact copy of min and max
	var minmax = map[string]int64{
		"min": min,
		"max": max,
	}
	// remove the negative sign if any exist
	for key, _ := range minmax {
		if minmax[key] < 0 {
			minmax[key] = minmax[key] * -1
		}
	}
	// create a variable to put the largest non negative value in from the "minmax map"
	var realmax int64
	if minmax["min"] > minmax["max"] {
		realmax = minmax["min"]
	} else {
		realmax = minmax["max"]
	}
	rand.Seed(time.Now().UnixNano())
	// use the realmax variable to attain a random value from rand.Int31n()
	value := rand.Int63n(realmax)
	// create a random chance that the return value is negative or positive
	if rand.Intn(2) == 1 {
		value = value * -1
	}
	// if we accidently create a value that is below the real minimum value(provided by the user)
	// redo the process until we get a value within the proper minimum and maximum bounds(provided by the user)
	for value < min || value > max {
		value = rand.Int63n(realmax)
		if rand.Intn(2) == 1 {
			value = value * -1
		}
	}
	return value
}
