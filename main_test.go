package main

import (
	"flag"
	"os"
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

func TestInitFlags(t *testing.T) {
	// Save original command line and reset flags
	originalArgs := os.Args
	originalCommandLine := flag.CommandLine

	defer func() {
		os.Args = originalArgs
		flag.CommandLine = originalCommandLine
	}()

	// Create a new flag set for this test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Test initFlags() function
	cfg := initFlags()

	// Test default values
	if cfg.help != false {
		t.Errorf("Expected help default to be false, got %v", cfg.help)
	}
	if cfg.version != false {
		t.Errorf("Expected version default to be false, got %v", cfg.version)
	}

	// Test that flags can be parsed
	testArgs := []string{
		"epoch.exe",
		"-?",
		"-v",
	}

	// Reset flag set and reinitialize
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cfg = initFlags()

	// Parse test arguments
	err := flag.CommandLine.Parse(testArgs[1:])
	if err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	// Verify flags were set correctly
	if !cfg.version {
		t.Error("Expected version flag to be true")
	}
}
