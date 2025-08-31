package main

import (
	"strings"
	"testing"
)

func TestHumanReadable(t *testing.T) {
	actual := epochToHumanReadable("1621258963", true)
	expected := "2021-05-17T13:42:43Z"

	compare(actual, expected, t)
}

func compare(actual, expected string, t *testing.T) {
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
