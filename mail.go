package rdg

import (
	"bytes"
	"strconv"
)

// Street returns a random street address
// e.g. "123 Foo St APT 2"
func Street() string {
	street := Street1()
	if Bool() {
		street += " " + Street2()
	}
	return street
}

// Street1 returns a random primary street address
// e.g. "123 Foo St"
func Street1() string {
	var street bytes.Buffer

	// number
	street.WriteString(strconv.FormatInt(int64(IntBetween(1, 99999)), 10))

	// prefix
	if Bool() {
		street.WriteString(" " + streetPrefix[IntBetween(0, len(streetPrefix)-1)])
	}

	// name
	street.WriteString(" " + streetNames[IntBetween(0, len(streetNames)-1)])

	// suffix
	if Bool() {
		street.WriteString(" " + streetSuffix[IntBetween(0, len(streetSuffix)-1)])
	}

	// type
	street.WriteString(" " + streetTypes[IntBetween(0, len(streetTypes)-1)])

	return street.String()
}

// Street2 returns a random secondary street address
// e.g. "Apt 2"
func Street2() string {
	var street bytes.Buffer
	street.WriteString(secondaryStreetTypes[IntBetween(0, len(secondaryStreetTypes)-1)])
	if Bool() {
		// letter
		street.WriteString(" " + alphabit[IntBetween(0, len(alphabit)-1)])
	} else {
		// number
		street.WriteString(" " + strconv.FormatInt(int64(IntBetween(1, 99)), 10))
	}
	return street.String()
}

// City returns a random city name
// e.g. "San Jose"
func City() string {
	return cities[IntBetween(0, len(cities)-1)]
}

// State returns a random State name
// e.g. "California"
func State() string {
	return states[IntBetween(0, len(states)-1)]
}

// Zip returns a random zip name
// e.g. "90001" <--> "99999"
func Zip() string {
	return strconv.FormatInt(int64(IntBetween(90001, 99999)), 10)
}

// Country returns a random country name
// e.g. "United States"
func Country() string {
	return countries[IntBetween(0, len(countries)-1)]
}
