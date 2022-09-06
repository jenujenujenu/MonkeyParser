package MonkeyParser

import (
	"os"
	"testing"
)

func TestFromFile(t *testing.T) {
	f, err := os.Open("test.md")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	FromFile(f)
}
