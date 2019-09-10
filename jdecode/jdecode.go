package jdecode

import (
	"fmt"
	"strings"
)

var family = []string{
	"QzF3n6/9CAtpu0O",
	"B1IREhcSyrleKvMW8LXx",
	"7N-dVbwsY2g4oaJZGUDj",
	"iHkq.mPf5T",
}

const magic = "$9$"

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

var numAlpha map[int]rune
var alphaNum map[rune]int

var extra map[rune]int

// Decode Junos password $9$
func Decode(passwordCoded string) (string, error) {
	initialize()
	chars := strings.TrimPrefix(passwordCoded, magic)
	var first string
	first, chars = nibble(chars, 1)
	firstR := []rune(first)
	_, chars = nibble(chars, extra[firstR[0]])
	prev := first
	passwordDecoded := ""
	for chars != "" {
		decode := encoding[len(passwordDecoded)%len(encoding)]
		var nib string
		nib, chars = nibble(chars, len(decode))
		var gaps []int
		for _, i := range nib {
			g := gap(prev, string(i))
			prev = string(i)
			gaps = append(gaps, g)
		}
		charsDecoded, _ := gapDecode(gaps, decode)
		/*if err != nil {
			return "", err
		}*/
		passwordDecoded += charsDecoded
	}
	return passwordDecoded, nil
}

// initialize fill numAlpha, alphaNum, extra with family
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

// gapDecode
func gapDecode(gaps []int, dec []int) (string, error) {
	var num int

	if len(gaps) != len(dec) {
		return "", fmt.Errorf("nibble and decode size not the same")
	}
	for i, gap := range gaps {
		num += gap * dec[i]
	}
	return string(num % 256), nil
}

func nibble(cref string, length int) (string, string) {
	crefSplit := strings.Split(cref, "")

	if len(crefSplit) < length {
		return strings.Join(crefSplit[0:], ""), ""
	}

	return strings.Join(crefSplit[0:length], ""), strings.Join(crefSplit[length:], "")
}

// gap betwean characters
func gap(c1 string, c2 string) int {
	c1rune := []rune(c1)
	c2rune := []rune(c2)

	return pmod((alphaNum[c2rune[0]]-alphaNum[c1rune[0]]), (len(numAlpha))) - 1
}

// modulus positive (same as python)
func pmod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
