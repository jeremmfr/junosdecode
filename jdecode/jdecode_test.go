package jdecode

import (
	"testing"
)

// TestDecode example password
func TestDecode(t *testing.T) {
	junwordCoded := "$9$1HFIyKXxdsgJ-VH.Pfn6lKMXdsZUi5Qnikfz"
	word := "testPassWord"

	passwordDecoded, err := Decode(junwordCoded)
	if err != nil {
		t.Errorf("error on decode %v", err)
	}
	if passwordDecoded != word {
		t.Errorf("decode password failed")
	}
}
