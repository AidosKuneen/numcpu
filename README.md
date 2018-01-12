[![Build Status](https://travis-ci.org/AidosKuneen/numcpu.svg?branch=master)](https://travis-ci.org/AidosKuneen/numcpu)
[![GoDoc](https://godoc.org/github.com/AidosKuneen/numcpu?status.svg)](https://godoc.org/github.com/AidosKuneen/numcpu)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/AidosKuneen/numcpu/master/LICENSE)


NumcCPU
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

