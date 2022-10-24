package junosdecode_test

import (
	"math/rand"
	"testing"

	"github.com/jeremmfr/junosdecode"
)

const (
	junWordCoded   = "$9$1HFIyKXxdsgJ-VH.Pfn6lKMXdsZUi5Qnikfz"
	junWordDecoded = "testPassWord"
)

// TestDecode decode example password.
func TestDecodePassword(t *testing.T) {
	passwordDecoded, err := junosdecode.Decode(junWordCoded)
	if err != nil {
		t.Errorf("error on decode %v", err)
	}
	if passwordDecoded != junWordDecoded {
		t.Errorf("decode password failed")
	}
}

// TestDecodeBadEncoded try decode encrypted secret but with missing or extra character.
func TestDecodeBadEncoded(t *testing.T) {
	// remove last character
	if passwordDecoded, err := junosdecode.Decode(junWordCoded[:len(junWordCoded)-1]); err == nil {
		t.Errorf("missing character in junWordCoded not detected, passwordDecoded: %s", passwordDecoded)
	}
	// add last character
	if passwordDecoded, err := junosdecode.Decode(junWordCoded + string(rune(rand.Intn(26)))); err == nil { //nolint: gosec
		t.Errorf("extra character in junWordCoded not detected, passwordDecoded: %s", passwordDecoded)
	}
	// add two last characters
	if passwordDecoded, err := junosdecode.Decode(
		junWordCoded + string(rune(rand.Intn(26))) + string(rune(rand.Intn(26)))); err == nil { //nolint: gosec
		t.Errorf("extra characters in junWordCoded not detected, passwordDecoded: %s", passwordDecoded)
	}
}
