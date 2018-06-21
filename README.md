# interview-prep
[![Go Report Card](https://goreportcard.com/badge/github.com/crazcalm/interview-prep)](https://goreportcard.com/report/github.com/crazcalm/interview-prep)  [![Build Status](https://api.travis-ci.org/crazcalm/interview-prep.svg?branch=master)](https://travis-ci.org/crazcalm/interview-prep)

This repository consists of the material that I wrote out (or looked up) while preparing for an interview.


## Layout
### Chapter X
Chapter X directories correspond to chapters in the Cracking the Coding Interview book.

### Data structures
That directory contains a bunch of data structures that I either wrote or found.

### Goroutines
This section consists of examples taken for The Go Programming Language book.

# Run Tests:
[Running benchmark test review](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

Run test in single directory:

	go test -run=. -bench=.

Run tests in all directories:

	go test ./... -run=. -bench=.
