# treecli
[![Build Status][travis-badge]][travis]

[travis-badge]: https://travis-ci.org/munjeli/treecli.svg?branch=master
[travis]: https://travis-ci.org/munjeli/treecli

Graph theory is a core skill for computer science yet I didn't have any tools that
allowed me to generate trees and graphs from columnar data and see the performance
of various implementations and algorithms. This cli is meant to give the user a
way to compare several ways of approaching a problem with finite data and easy to
use commands.
## Description
Treecli is a Command Line Interface for exploring graph theory from columnar data.
You can feed a data file to the cli and it will create a binary tree. You can insert
a node into a generated tree with the insert command.
## Usage

## Install
To install, use `go get`:

```bash
$ go get -d github.com/munjeli/treecli
```
cd into the directory and run `make` to run gofmt and build a single binary with [gox](https://github.com/mitchellh/gox).
There isn't an official distribution of treecli yet, so the Makefile is just set up for
development.

## Contribution

1. Fork ([https://github.com/munjeli/treecli/fork](https://github.com/munjeli/treecli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run the gofmt and the test suite with the `make` command and confirm that it passes
1. Create a new Pull Request

## Author

[munjeli](https://github.com/munjeli)
