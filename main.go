package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	binary = `epoch`

	usage = "\n" + binary + " prints the local time corresponding to the given Unix time\n" +
		"given as argument (in seconds since January 1, 1970 UTC).\n\n" +
		"Usage: %s [OPTIONS] <int64> <int64>...\n\nOPTIONS:\n"
)

// set via -ldflags
var version string

func main() {
	parseFlags()

	for _, arg := range flag.Args() {
		fmt.Println(epochToHumanReadable(arg))
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
