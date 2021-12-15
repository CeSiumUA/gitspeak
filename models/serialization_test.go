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

func TestRepositoriesDeserialization(t *testing.T) {
	expectedLength := 100
	text, err := os.ReadFile("repos_test.json")
	check(err)
	reposFromBytes, err := DeserializeArrayBytes(text)
	check(err)
	reposFromBytesLen := len(reposFromBytes)
	if reposFromBytesLen != expectedLength {
		t.Errorf("Got invalid repos count, %d expected, got %d", expectedLength, reposFromBytesLen)
	}
	reposFromStr, err := DeserializeArrayString(string(text))
	check(err)
	reposFromStrLen := len(reposFromStr)
	if reposFromStrLen != expectedLength {
		t.Errorf("Got invalid repos count, %d expected, got %d", expectedLength, reposFromStrLen)
	}
}

func TestRepoLanguagesDeserialization(t *testing.T) {
	languages := map[string]int64{
		"Assembly": 1811,
		"Rust":     192,
	}
	text, err := os.ReadFile("langs_test.json")
	check(err)
	langsFromBytes, err := DeserializeLanguagesFromBytes(text)
	check(err)
	langsLength := len(*langsFromBytes)
	if langsLength != len(languages) {
		t.Errorf("Got %d elements deserialized, expected %d", langsLength, len(languages))
	}
	for key, value := range *langsFromBytes {
		if value != languages[key] {
			t.Errorf("Values don't match! Got from deserializer: %d, expected: %d", value, languages[key])
		}
	}
}
