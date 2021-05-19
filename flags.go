package main

import (
	"flag"
	"fmt"
	"os"
)

var utc bool
var showVersion bool

func parseFlags() {
	flag.BoolVar(&utc, "utc", false, "print time as UTC (Coordinated Universal Time)")
	flag.BoolVar(&showVersion, "version", false, "print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, binary)
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	} else if showVersion || flag.Arg(0) == "version" {
		fmt.Printf("%s version %s\n", binary, version)
		os.Exit(0)
	}
}
