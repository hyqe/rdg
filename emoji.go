package rdg

func Emoji() string {
	return emojis[IntBetween(0, len(emojis))]
}
