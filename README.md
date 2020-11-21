# Random Data Generator (rdg) 

Generate Random Data for mocking and testing purposes.


## TODO:

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
func PhoneNumberUS() string 
```

### Business

```go
func CompanyName() string 
func JobTitle() string 
func Department() string 
```

### Time

```go
func Time(min, max time.Time) time.Time 
func Date() string 
func DateTime() string 
```


### HTTP URL
```go
func URL() url.Url 
func DomainName() string 
func URLQueryString() string 
func URLSchema() string 
```

### Money

```go
func Money(min, man float64) string 
func Dollar(min, max int64) string
```

### Emoji

```go
func Emoji() string
```

### Data Structures

```go
func FieldName() string 
func JSON() string 
func XML() string 
func HTML() string
```

### Text

```go
func Word() string
func Sentence(words uint64) string
func Paragraph(sentences uint64) string
```

### Utilities

Pick a random number between min and max
```go
// func Number(min, max int64) int64
func Number(1,10)
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
