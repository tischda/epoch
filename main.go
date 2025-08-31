package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

var (
	flagHelp    = flag.Bool("help", false, "displays this help message")
	flagUtc     = flag.Bool("utc", false, "print time as UTC (Coordinated Universal Time)")
	flagVersion = flag.Bool("version", false, "print version and exit")
)

func init() {
	flag.BoolVar(flagHelp, "h", false, "")
	flag.BoolVar(flagVersion, "v", false, "")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] <int64> <int64>...

Prints the local time corresponding to the given Unix time
given as argument (in seconds since January 1, 1970 UTC).

OPTIONS:

  -h, --help
        display this help message
  -utc
        print time as UTC (Coordinated Universal Time)
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` -utc 1521258963 1621258987
    2018-03-17T03:56:03Z
    2021-05-17T13:43:07Z`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || *flagVersion {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if *flagHelp {
		flag.Usage()
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	for _, arg := range flag.Args() {
		fmt.Println(epochToHumanReadable(arg, *flagUtc))
	}
}

func epochToHumanReadable(s string, utc bool) string {
	epoch, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
	t := time.Unix(epoch, 0)
	if utc {
		t = t.UTC()
	}
	return t.Format(time.RFC3339)
}
