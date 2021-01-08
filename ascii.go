package rdg

func Ascii() string {
	return ascii_emoticons[IntBetween(0, len(ascii_emoticons))]
}
