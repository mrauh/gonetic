package gonetic

import (
	"testing"
)

type phonetic struct {
	str      string
	phonCode string
}

var phonetics = []phonetic{
	phonetic{str: "Müller-Lüdenscheidt", phonCode: "65752682"},
	phonetic{str: "Wikipedia", phonCode: "3412"},
	phonetic{str: "Breschnew", phonCode: "17863"},
}

//--------------------------------------------------------------------------

func TestPhonetics(t *testing.T) {
	for _, p := range phonetics {
		code := NewPhoneticCode(p.str)

		if code != p.phonCode {
			t.Errorf("Error: expected %v, got %v",
				p.phonCode, code)
		}
	}
}

//--------------------------------------------------------------------------

type duplicate struct {
	strFull     string
	strExpected string
}

//--------------------------------------------------------------------------

var duplicates = []duplicate{
	duplicate{strFull: "65575268822", strExpected: "65752682"},
	duplicate{strFull: "6557526882", strExpected: "65752682"},
	duplicate{strFull: "1234567899", strExpected: "123456789"},
	duplicate{strFull: "1123456789", strExpected: "123456789"},
	duplicate{strFull: "1111111111", strExpected: "1"},
}

//--------------------------------------------------------------------------

func TestRemoveDuplicates(t *testing.T) {
	for _, dup := range duplicates {
		strRem := removeDuplicates(dup.strFull)

		if strRem != dup.strExpected {
			t.Errorf("Error: expected %v, got %v",
				dup.strExpected, strRem)
		}
	}
}

//--------------------------------------------------------------------------
