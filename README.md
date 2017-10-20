# fedACH 
[![Build Status](https://travis-ci.org/thylong/fedach.svg?branch=master)](https://travis-ci.org/thylong/fedach) [![Go Report Card](https://goreportcard.com/badge/github.com/thylong/fedach)](https://goreportcard.com/report/github.com/thylong/fedach) [![GoDoc](https://godoc.org/github.com/thylong/fedach?status.png)](https://godoc.org/github.com/thylong/fedach)

This Go package allows you to Marshal or Unmarshal from/to [FedACH Participant RDFIs With Commercial Receipt Volume](https://www.frbservices.org/operations/epayments/epayments.html).

You can rely on this package to have the latest official version of the Fed file.

## Installation

```bash
go get -u github.com/thylong/fedach
```

## Examples

### Get current Fed file version

```go
package main

import (
	"fmt"

	"github.com/thylong/fedach"
)

func main() {
	current := fedach.GetCurrentDirectoryFile()

	fmt.Printf("%#v\n", current)
}
```

### Unmarshal from file

```go
package main

import (
	"fmt"

	"github.com/thylong/fedach"
)

func main() {
    fileContent, err := ioutil.ReadFile("sample/FedACHdir.txt")
	if err != nil {
		panic(err)
	}

	var routingDirectoryFile []fedach.RoutingDirectoryRecord
	if err := fedach.Unmarshal(fileContent, &routingDirectoryFile); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", routingDirectoryFile)
}
```

### Marshal to file

#### Using provided data structure

```go
package main

import (
	"fmt"

	"github.com/thylong/fedach"
)

func main() {
	routingDirectoryFile := []fedach.RoutingDirectoryRecord{
		{
			"011000015",
			"O",
			"011000015",
			"0",
			"122415",
			"000000000",
			"FEDERAL RESERVE BANK                ",
			"1000 PEACHTREE ST N.E.              ",
			"ATLANTA             ",
			"GA",
			"30309",
			"4470",
			"877",
			"372",
			"2457",
			"1",
			"1",
			"     ",
		},
	}
	encoded, err := fedach.Marshal(routingDirectoryFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", encoded)
}
```

#### Using slice

```go
package main

import (
	"fmt"

	"github.com/thylong/fedach"
)

func main() {
	routingDirectoryFile := [][]string{
		{
			"011000015",
			"O",
			"011000015",
			"0",
			"122415",
			"000000000",
			"FEDERAL RESERVE BANK                ",
			"1000 PEACHTREE ST N.E.              ",
			"ATLANTA             ",
			"GA",
			"30309",
			"4470",
			"877",
			"372",
			"2457",
			"1",
			"1",
			"     ",
		},
	}
	encoded, err := fedach.Marshal(routingDirectoryFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", encoded)
}
```

## Notes

Please make sure you read the frbservices agreement:
https://www.frbservices.org/EPaymentsDirectory/agreement.html
