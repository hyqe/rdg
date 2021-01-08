package rdg

import (
	"testing"
)

func run(check string, array []string) bool {
	var length = len(array)
	for x := 0; x < length; x++ {
		if check == array[x] {
			return true
		}
	}
	return false
}

func TestEmoji(t *testing.T) {

	//var value string
	for i := 0; i < 1000; i++ {
		if !run(Emoji(), emojis) {
			t.Fatalf("String returned from Emoji() does not exist in emojis-array\n")
		}
		if !run(Ascii(), ascii_emoticons) {
			t.Fatalf("String returned from ASCIIArt() does not exist in ascIIArts-array\n")
		}
	}
}
