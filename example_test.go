package junosdecode_test

import (
	"fmt"

	"github.com/jeremmfr/junosdecode"
)

func Example() {
	junWordCoded := "$9$1HFIyKXxdsgJ-VH.Pfn6lKMXdsZUi5Qnikfz"
	if passwordDecoded, err := junosdecode.Decode(junWordCoded); err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(passwordDecoded)
		// Output: testPassWord
	}
}
