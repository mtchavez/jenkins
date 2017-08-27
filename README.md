# Jenkins Hash

[![Latest Version](http://img.shields.io/github/release/mtchavez/jenkins.svg?style=flat-square)](https://github.com/mtchavez/jenkins/releases)
[![Build Status](https://travis-ci.org/mtchavez/jenkins.svg?branch=master)](https://travis-ci.org/mtchavez/jenkins)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/jenkins)
[![Coverage Status](https://coveralls.io/repos/github/mtchavez/jenkins/badge.svg)](https://coveralls.io/github/mtchavez/jenkins)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/jenkins)](https://goreportcard.com/report/github.com/mtchavez/jenkins)

Golang Jenkins hash

## Install

`go get -u github.com/mtchavez/jenkins`

## Usage

Jenkins conforms to the [Hash32](http://golang.org/pkg/hash/#Hash32) interface from the Go standard library

```go
// Create a new hash
jenkhash := New()

// Write a string of bytes to hash
key := []byte("my-random-key")
length, err := jenkhash(key)

// Get uint32 sum of hash
sum := jenkhash.Sum32()

// Sum hash with byte string
sumbytes := jenkhash.Sum(key)
```

## Testing

Uses [Ginkgo][ginkgo]  and [Gomega][gomega] for testing.

Run via `make test` which will run `go test -cover`

## Documentation

Docs on [godoc](http://godoc.org/github.com/mtchavez/jenkins)

[ginkgo]: https://github.com/onsi/ginkgo
[gomega]: https://github.com/onsi/gomega
