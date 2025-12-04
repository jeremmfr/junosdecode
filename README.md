# junosdecode

[![Go Reference](https://pkg.go.dev/badge/github.com/jeremmfr/junosdecode.svg)](https://pkg.go.dev/github.com/jeremmfr/junosdecode)
[![Release](https://img.shields.io/github/v/release/jeremmfr/junosdecode)](https://github.com/jeremmfr/junosdecode/releases)
[![Report Card](https://goreportcard.com/badge/github.com/jeremmfr/junosdecode)](https://goreportcard.com/report/github.com/jeremmfr/junosdecode)
[![Go Status](https://github.com/jeremmfr/junosdecode/workflows/Go%20Tests/badge.svg)](https://github.com/jeremmfr/junosdecode/actions)
[![Lint Status](https://github.com/jeremmfr/junosdecode/workflows/GolangCI-Lint/badge.svg)](https://github.com/jeremmfr/junosdecode/actions)
[![codecov](https://codecov.io/gh/jeremmfr/junosdecode/branch/main/graph/badge.svg)](https://codecov.io/gh/jeremmfr/junosdecode)

Library to decode Juniper secret hashes (`$9$`)

Based on [Crypt-Juniper](http://search.cpan.org/dist/Crypt-Juniper/lib/Crypt/Juniper.pm) and [taktv6/junoscrypt](https://github.com/taktv6/junoscrypt)

## Usage

```go
package main

import (
     "fmt"
     "github.com/jeremmfr/junosdecode"
)

func main() {
     junWordCoded := "$9$1HFIyKXxdsgJ-VH.Pfn6lKMXdsZUi5Qnikfz"
     if passwordDecoded, err := junosdecode.Decode(junWordCoded); err != nil {
          fmt.Print(err.Error())
     } else {
          fmt.Print(passwordDecoded)
     }
}
```

Play : [here](https://play.golang.org/p/tibSxKFM_zx)
