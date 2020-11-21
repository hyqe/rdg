# Random Data Generator (rdg) 

Generate Random Data for mocking and testing purposes. This Module will focuses manly the random generation of data and will try to avoid richly formatting data whenever possible. This module will be designed to pair nicely with the [github.com/hyqe/rich](https://github.com/hyqe/rich) module.


## TODOs:

### Names

```go
func Prefix() string 
func FirstName() string 
func LastName() string 
func Suffix() string 
```

### Gender Names
```go
func PrefixMale() string 
func PrefixFemale() string 
func FirstNameMale() string 
func FirstNameFemale() string 
func SuffixMale() string 
func SuffixFemale() string 
```

### Mail Address

```go
func Street() string 
func City() string 
func State() string 
func Zip() string 
func Country() string 
```

### Email Address

```go
func EmailAddress() string 
func EmailAddressMale() string 
func EmailAddressFemale() string 
func EmailAddressResource() string 
```

### Phone Number

```go
func PhoneNumber() string 
```

### Business

```go
func CompanyName() string 
func JobTitle() string 
func Department() string 
```

### Time

```go
func Time() time.Time 
func TimeBetween(min, max time.Time)
```


### URL
```go
func URL() url.URL
func DomainName() string 
func URLQuery() string 
func URLSchema() string
func URLPath() string
func URLFragment() string
```

### Money

```go
func Money() string 
func MoneyBetween(min, man float64) string 
```

### Emoji

```go
func Emoji() string
func ASCIIArt() string
```

### Text

```go
func Word() string
func Sentence(words uint64) string
func Paragraph(sentences uint64) string
```


### Data Structures

```go
func FieldName() string
func JSON() string 
func XML() string 
func HTML() string
```

### SQL

```go
func SQL() string
func SQLSelect() string
func SQLUpdate() string
func SQLInsert() string
func SQLDelete() string
```

### Utilities

Pick a random number between min and max
```go
func Number() uint64
func NumberBetween(min, max uint64) uint64
```

Pick a random string  from a list of string s
```go
// func String(list ...string ) string 
func String("Data", "To", "Pick", "From") string 
```

Get a random boolean
```go
// func Bool() bool
func Bool() string 
```

Pick a random value from a list
```go
// func Any(list ...interface) interface
func Any("Foo", 'b', 30, 7.1) string 
```


Format a string with list of random formatters
```go
type Formatter = func(string) string
func Format(string, ...Formatter) string
```