package rdg

// Formatter is any function that can format a string
// Formatters should never panic.
type Formatter = func(string) string

// Format randomly formats a string with a list of formatters
func Format(string, ...Formatter) string {
	return ""
}
