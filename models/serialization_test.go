package models

import (
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestDeserialization(t *testing.T) {
	expectedLength := 100
	text, err := os.ReadFile("test.json")
	check(err)
	reposFromBytes := DeserializeArrayBytes(text)
	reposFromBytesLen := len(reposFromBytes)
	if reposFromBytesLen != expectedLength {
		t.Errorf("Got invalid repos count, %d expected, got %d", expectedLength, reposFromBytesLen)
	}
	reposFromStr := DeserializeArrayString(string(text))
	reposFromStrLen := len(reposFromStr)
	if reposFromStrLen != expectedLength {
		t.Errorf("Got invalid repos count, %d expected, got %d", expectedLength, reposFromStrLen)
	}
}
