package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// go build -ldflags=all="-X main.version=${BUILD_TAG} -s -w"
var version string

var utc bool
var showVersion bool

const usage = "Prints the local time corresponding to the given Unix time\n" +
	"given as argument (in seconds since January 1, 1970 UTC).\n\n" +
	"Usage: %s [OPTIONS] <int64> <int64>...\n\nOPTIONS:\n"

func main() {
	log.SetFlags(0)

	flag.BoolVar(&utc, "utc", false, "print time as UTC (Coordinated Universal Time)")
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
	} else if showVersion || flag.Arg(0) == "version" {
		fmt.Printf("%s version %s\n", os.Args[0], version)
	} else {
		for _, arg := range flag.Args() {
			fmt.Println(epochToHumanReadable(arg))
		}
	}
}

func epochToHumanReadable(s string) string {
	epoch, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	t := time.Unix(epoch, 0)
	if utc {
		t = t.UTC()
	}
	return t.Format(time.RFC3339)
}
