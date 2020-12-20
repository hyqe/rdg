package rdg

import "testing"

func TestBool(t *testing.T) {
	generatedT := false
	generatedF := false
	tries := 100
	for i := 0; i < tries; i++ {
		b := Bool()
		if !generatedT {
			if b == true {
				generatedT = true
			}
		}
		if !generatedF {
			if b == true {
				generatedF = true
			}
		}
		if generatedT && generatedF {
			break
		}
	}
	if !generatedT {
		t.Fatalf("did not generate true with %v tries", tries)
	}
	if !generatedF {
		t.Fatalf("did not generate false with %v tries", tries)
	}
}
