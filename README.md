# Dictionary Dash

An engineering exercise implemented in Go.

Given a dictionary and a start and end word, this program finds the minimum number of single-letter transformations required to transform the start word into the end word, where each intermediate word is also in the dictionary.

This entire project can be cloned directly from github via:
https://github.com/nerophon/dictdash

## Prerequisites

1. The [__Go Programming Langugage__][0] must be installed to build, test, and install this software.

## Installation

1. Clone this project.
2. `cd` to the project directory
3. run `go install`

The software will be installed to the `$GOPATH/bin` directory by default.

## Testing & Benchmarking

This software includes extensive unit tests. They can be run as per standard for Go tests:

1. `cd` to the source folder with test files you wish to run
2. run `go test`

Benchmarks exist for key steps in the process. These can be run from the root project directory, via the `dictdash_test.go` file. To run them, use the following standard Go benchmarking command:

```
go test -bench=.
```

The scanning of the 200k+ words sample dictionary takes several seconds, so if you wish to get a statistically significant benchmark for it please specify the Scan function and use the `-benchtime` flag with a time greater than 1 second, for instance:

```
go test -bench=BenchmarkScan -benchtime=20s
```

Please be aware that this kind of benchmark can put strain on a computer's CPU resources and should be run with care.

## Launching

1. `cd` to the install directory, usually `$GOPATH/bin`
2. run `./dictdash`

## Operation

It is recommended that users copy the sample dictionary `dict.txt` found in the root project directory into the `$GOPATH/bin` directory (or wherever the executable is located if elsewhere).

The `scan` command can then be used without specifying a path argument:

```
> scan
```

Scanning the sample dictionary will take a few seconds, as the program builds a graph of more than 200k English words.

Once scanning is complete, word count and frequency by letter count will be displayed. It will be possible to search the graph for transform paths using the `search` command. For example:

```
> search bounce lather
```

This produces the following output:

```
The shortest path between bounce and lather is 29 transformations long.
Full path:
0: bounce
1: jounce
2: jaunce
3: launce
4: launch
5: caunch
6: clunch
7: clutch
8: clitch
9: slitch
10: stitch
11: stetch
12: stench
13: stanch
14: starch
15: sparch
16: sparth
17: swarth
18: swarty
19: starty
20: starvy
21: starve
22: staree
23: starer
24: searer
25: seater
26: setter
27: letter
28: latter
29: lather
```

## Custom Dictionaries

This exercise has been implemented with several assumptions in mind concerning acceptable format for the scanned dictionary. To guarantee reasonable performance, validation of this format is not performed by this application; therefore it is up to the user to ensure the dictionary is properly formatted if acceptable results from this application are desired.

The dictionary should be formatted as a whitespace-delimited list of words containing only lowercase latin utf8 characters from a to z.

In this case, whitespace is defined as per the Go library function `unicode.IsSpace()`, described here:
https://golang.org/pkg/unicode/#IsSpace

## Performance

Emphasis was placed upon getting good performance from the `Search` functionality, under the assumption that the dictionary would not need to be reloaded often. Due to time constraints, concurrency was only used in the most obvious locations (`grapher.link()` and `grapher.compress()`). I believe it is possible to improve the performance of both the `grapher.scan()` and `search.path()` functions using concurrency, but doing so would be non-trivial and require further research.


[0]: https://golang.org/dl/
