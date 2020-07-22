# junosdecode

[![GoDoc](https://godoc.org/github.com/jeremmfr/junosdecode?status.svg)](https://godoc.org/github.com/jeremmfr/junosdecode)
[![Report Card](https://goreportcard.com/badge/github.com/jeremmfr/junosdecode)](https://goreportcard.com/report/github.com/jeremmfr/junosdecode)
[![Go Status](https://github.com/jeremmfr/junosdecode/workflows/Go%20Tests/badge.svg)](https://github.com/jeremmfr/junosdecode/actions)
[![Lint Status](https://github.com/jeremmfr/junosdecode/workflows/GolangCI-Lint/badge.svg)](https://github.com/jeremmfr/junosdecode/actions)
[![codecov](https://codecov.io/gh/jeremmfr/junosdecode/branch/master/graph/badge.svg)](https://codecov.io/gh/jeremmfr/junosdecode)

Library to decode Juniper password hashes ($9$)

Based on http://search.cpan.org/dist/Crypt-Juniper/lib/Crypt/Juniper.pm and https://github.com/taktv6/junoscrypt

## Usage

```
package main

import (
     "fmt"
     jdecode "github.com/jeremmfr/junosdecode"
)

func main() {
     junwordCoded := "$9$1HFIyKXxdsgJ-VH.Pfn6lKMXdsZUi5Qnikfz"
     passwordDecoded, err := jdecode.Decode(junwordCoded)
     if err != nil {
          fmt.Print(err.Error())
     }
     
     fmt.Print(passwordDecoded)
}

```
Play : https://play.golang.org/p/HpGiCYMjV5W
