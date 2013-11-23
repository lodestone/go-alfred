package Alfred

import (
	// "fmt"
	"testing"
)

func TestSanity(t *testing.T) {
	ga := NewAlfred("daftar")
	ga.SetNoResultTxt("Test Number 1!")
	ga.WriteToAlfred()
}
