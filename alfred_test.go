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

func TestBasics(t *testing.T) {
	tests := []struct {
		id       string
		expected string
	}{
		id:       "TestBasic",
		expected: "<>",
	}

	for _, test := range tests {
		ga := NewAlfred(test.id)
		res, err := ga.XML()
		if err != nil {
			t.Fatalf("%s has faild with: %v", test.id, err)
		}
		if res != test.expected {
			t.Errorf("Expected %v but received %v\n", test.expected, res)
		}
	}
	ga := NewAlfred("TestBasic")
	res, err := ga.XML()
}
