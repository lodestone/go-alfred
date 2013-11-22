package alfred

import (
	"fmt"
	"testing"
)

func TestSanity(t *testing.T) {
	ga = Alfred.NewAlfred()
	ga.SetNoResultTxt("Test Number 1!")
}
