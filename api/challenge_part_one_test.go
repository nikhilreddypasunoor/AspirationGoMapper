package api

import (
	"testing"
)

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	type args struct {
		input, expectedResult string
	}

	var testcases = []args{
		args{
			input:          "Aspiration.com",
			expectedResult: "asPirAtiOn.cOm",
		},
		args{
			input:          "",
			expectedResult: "",
		},
		args{
			input:          "ASPIRATION",
			expectedResult: "asPirAtiOn",
		},
		args{
			input:          "aspiration",
			expectedResult: "asPirAtiOn",
		},
		args{
			input:          "ASPIRATION...COM",
			expectedResult: "asPirAtiOn...cOm",
		},
		args{
			input:          "    ",
			expectedResult: "    ",
		},
	}
	for _, test := range testcases {
		if output := CapitalizeEveryThirdAlphanumericChar(test.input); output != test.expectedResult {
			t.Errorf("Output %s not equal to expected %s", output, test.expectedResult)
		}
	}

}
