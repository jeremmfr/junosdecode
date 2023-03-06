package junosdecode_test

import (
	"errors"
	"fmt"
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
	t.Parallel()

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
	t.Parallel()

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

func TestDecodeEmptyEncoded(t *testing.T) {
	t.Parallel()

	// empty input
	if passwordDecoded, err := junosdecode.Decode(""); err == nil {
		t.Errorf("empty input not detected, passwordDecoded: %s", passwordDecoded)
	} else if !errors.Is(err, junosdecode.ErrEmptySecret) {
		t.Errorf("got unexpected error: %s", err)
	}
}

func TestDecodeNotEnoughCharsEncoded(t *testing.T) {
	t.Parallel()

	// MagicPrefix input
	if passwordDecoded, err := junosdecode.Decode(junosdecode.MagicPrefix); err == nil {
		t.Errorf("MagicPrefix input not detected, passwordDecoded: %s", passwordDecoded)
	} else if !errors.Is(err, junosdecode.ErrNotEnoughChars) {
		t.Errorf("got unexpected error: %s", err)
	}

	// MagicPrefix and 2 chars input
	if passwordDecoded, err := junosdecode.Decode(junosdecode.MagicPrefix + "!!"); err == nil {
		t.Errorf("MagicPrefix + 2 chars input not detected, passwordDecoded: %s", passwordDecoded)
	} else if !errors.Is(err, junosdecode.ErrNotEnoughChars) {
		t.Errorf("got unexpected error: %s", err)
	}
}

func TestDecodeParrallel(t *testing.T) {
	t.Parallel()

	for i, v := range []string{junWordCoded, junWordCoded, junWordCoded} {
		word := v
		t.Run(fmt.Sprintf("Loop number %d", i), func(t *testing.T) {
			t.Parallel()

			output, err := junosdecode.Decode(word)
			if err != nil {
				t.Errorf("got unexpected error: %s", err)
			}
			if output != junWordDecoded {
				t.Errorf("got unexpected decoded word: %s", output)
			}
		})
	}
}
