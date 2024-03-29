# changelog

## v1.1.1

BUG FIXES:

* fix panic when call `Decode()` in parallel
* fix panic when string input of `Decode()` is empty
* return error when there isn't enough characters in string input of `Decode()`

## v1.1.0

NOTES:

Go 1.15 or higher is now required to use this package

ENHANCEMENTS:

* export constant `MagicPrefix`
* export variable `ErrDiffGapDec`
* generate an error when missing or extra character(s) detected in encoded string
* add an example for [godoc](https://pkg.go.dev/github.com/jeremmfr/junosdecode#example_)

BUG FIXES:

* fix conversion of the character in string from the rune number (compatibility with latest golang version)
* fix name of Decode function variables (password to secret)

## v1.0.0

BUG FIXES:

* fix module path for ROOT
* fix linter error
* exclude go test of package

## v0.1.0

First release
