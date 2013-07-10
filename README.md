# Kölner Phonetik (Cologne Phonetic) implemented in Go

Gonetic implements the Kölner Phonetik (Cologne Phonetic) algorithm in Go. It 
is a translation of the php implementation of 
[deezaster](https://github.com/deezaster/germanphonetic) to Go.

[![Build Status](https://drone.io/github.com/mrauh/gonetic/status.png)](https://drone.io/github.com/mrauh/gonetic/latest)

## Installation

With [Go](http://www.golang.org) installed on your machine:

	$ go get github.com/mrauh/gonetic

## Usage / Example

```go
import (
	"github.com/mrauh/gonetic"
)

func main() {
	code := gonetic.NewPhoneticCode("Müller-Lüdenscheidt")
	println(code)
}
```
