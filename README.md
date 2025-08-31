[![Build Status](https://github.com/tischda/epoch/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/epoch/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/epoch/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/epoch/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tischda/epoch)](https://goreportcard.com/report/github.com/tischda/epoch)

# epoch

Prints the local time corresponding to the given Unix time given as argument (in seconds since January 1, 1970 UTC).

Inspired from https://www.epochconverter.com/

### Install

~~~
go install github.com/tischda/epoch@latest
~~~

### Usage

~~~
Usage: epoch [OPTIONS] <int64> <int64>...

OPTIONS:
  -utc
        print time as UTC (Coordinated Universal Time)
  -version
        print version and exit
~~~

### Examples

~~~
$ epoch 1621258987
2021-05-17T15:43:07+02:00

$ epoch -utc 1521258963 1621258987
2018-03-17T03:56:03Z
2021-05-17T13:43:07Z
~~~

The time layout is RFC3339.
