![Golangci lint](https://github.com/AidosKuneen/numcpu/actions/workflows/golangci-lint.yml/badge.svg)
![Tests](https://github.com/AidosKuneen/numcpu/actions/workflows/tests.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/AidosKuneen/numcpu/badge.svg?branch=master)](https://coveralls.io/github/AidosKuneen/numcpu?branch=master)
[![GoDoc](https://godoc.org/github.com/AidosKuneen/numcpu?status.svg)](https://godoc.org/github.com/AidosKuneen/numcpu)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/AidosKuneen/numcpu/master/LICENSE)


NumCPU
=====

## Overview

numcpu.NumCPU() returns the number of logical CPUs usable by the current process.

Android doesn't return max number of CPUs by runtime.NumCPU(),
but just returns number of "working" CPUs.
So this library checks sysfs directory directly for linux, and just returns runtime.NumCPU() for others.

## Requirements

* git
* go

are required to compile.


## Install
    $ go get -u github.com/AidosKuneen/numcpu


## Usage

```go
	import "github.com/AidosKuneen/numcpu"
	n:=numcpu.NumCPU()
```


## Dependencies and Licenses

```
Golang Standard Library       BSD 3-clause License
```
