package junosdecode

import (
	"errors"
	"strings"
)

// MagicPrefix: Junos encrypted secret prefix.
const MagicPrefix = "$9$"

// Decode Junos encrypted secret ($9$).
func Decode(encryptedSecret string) (secretDecoded string, _ error) {
	initialize()
	chars := strings.TrimPrefix(encryptedSecret, MagicPrefix)
	var first string
	first, chars = nibble(chars, 1)
	firstR := []rune(first)
	_, chars = nibble(chars, extra[firstR[0]])
	prev := first
	for chars != "" {
		decode := encoding[len(secretDecoded)%len(encoding)]
		var nib string
		nib, chars = nibble(chars, len(decode))
		var gaps []int
		for _, i := range nib {
			g := gap(prev, string(i))
			prev = string(i)
			gaps = append(gaps, g)
		}
		charsDecoded, err := gapDecode(gaps, decode)
		if err != nil {
			return "", err
		}
		secretDecoded += charsDecoded
	}

	return secretDecoded, nil
}

var family = []string{
	"QzF3n6/9CAtpu0O",
	"B1IREhcSyrleKvMW8LXx",
	"7N-dVbwsY2g4oaJZGUDj",
	"iHkq.mPf5T",
}

var encoding = [][]int{
	{
		1, 4, 32,
	},
	{
		1, 16, 32,
	},
	{
		1, 8, 32,
	},
	{
		1, 64,
	},
	{
		1, 32,
	},
	{
		1, 4, 16, 128,
	},
	{
		1, 32, 64,
	},
}

var (
	numAlpha map[int]rune
	alphaNum map[rune]int
	extra    map[rune]int
)

// initialize fill numAlpha, alphaNum, extra with family.
func initialize() {
	numAlpha = make(map[int]rune)
	alphaNum = make(map[rune]int)
	extra = make(map[rune]int)
	for i, r := range []rune(strings.Join(family, "")) {
		numAlpha[i] = r
		alphaNum[r] = i
	}

	for i, fam := range family {
		for _, c := range fam {
			extra[c] = 3 - i
		}
	}
}

// ErrDiffGapDec is returned when can't decode a character due to a missing or extra character(s).
//
// This is detected by a difference in length of the internal `gaps` and `decode` list.
var ErrDiffGapDec = errors.New("junosdecode: missing or extra character(s) (gaps and decode size not the same)")

// gapDecode.
func gapDecode(gaps []int, dec []int) (string, error) {
	var num int

	if len(gaps) != len(dec) {
		return "", ErrDiffGapDec
	}
	for i, gap := range gaps {
		num += gap * dec[i]
	}

	return string(rune(num % 256)), nil
}

func nibble(cref string, length int) (string, string) {
	crefSplit := strings.Split(cref, "")

	if len(crefSplit) < length {
		return strings.Join(crefSplit[0:], ""), ""
	}

	return strings.Join(crefSplit[0:length], ""), strings.Join(crefSplit[length:], "")
}

// gap betwean characters.
func gap(c1 string, c2 string) int {
	c1rune := []rune(c1)
	c2rune := []rune(c2)

	return pmod((alphaNum[c2rune[0]]-alphaNum[c1rune[0]]), (len(numAlpha))) - 1
}

// modulus positive (same as python).
func pmod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}
