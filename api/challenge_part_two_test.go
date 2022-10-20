package api

import (
	"testing"
)

func TestMapString(t *testing.T) {
	type args struct {
		input          string
		expectedResult string
		idx int
	}

	var testcases = []args{
		args{
			input:          "Aspiration.com",
			expectedResult: "asPirAtiOn.cOm",
			idx: 3,
		},
		args{
			input:          "",
			expectedResult: "",
			idx: 3,
		},
		args{
			input:          "ASPIRATION",
			expectedResult: "asPirAtiOn",
			idx: 3,
		},
		args{
			input:          "aspiration",
			expectedResult: "asPirAtiOn",
			idx: 3,
		},
		args{
			input:          "ASPIRATION...COM",
			expectedResult: "asPirAtiOn...cOm",
			idx: 3,
		},
		args{
			input:          "    ",
			expectedResult: "    ",
			idx: 3,
		},
		args{
			input:          "Aspiration.com",
			expectedResult: "aspiRatioN.com",
			idx: 5,
		},
	}
	for _, test := range testcases {
		s := ModString{
			idx:      test.idx,
			inputStr: test.input,
		}
		if MapString(&s); s.resStr != test.expectedResult {
			t.Errorf("Output %s not equal to expected %s", s.resStr, test.expectedResult)
		}
	}

}
