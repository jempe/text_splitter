package splitter

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitTextInChunks(t *testing.T) {
	testCases := []struct {
		name           string
		text           string
		chunkLength    int
		rangeLength    int
		delimiters     []rune
		expectedResult []string
	}{
		{
			name:           "Split by whitespace",
			text:           "This is a test string to split into chunks",
			chunkLength:    10,
			rangeLength:    5,
			delimiters:     []rune{' '},
			expectedResult: []string{"This is a test", " string to", " split into", " chunks"},
		},
		{
			name:           "Split by comma",
			text:           "apple,banana,cherry,date,elderberry",
			chunkLength:    5,
			rangeLength:    3,
			delimiters:     []rune{','},
			expectedResult: []string{"apple", ",banana", ",cherry", ",date", ",elderbe", "rry"},
		},
		{
			name:           "Split by multiple delimiters",
			text:           "This, is. a! test; string: to- split_ into+ chunks",
			chunkLength:    5,
			rangeLength:    3,
			delimiters:     []rune{',', '.', '!', ';', ':', '-', '_', '+'},
			expectedResult: []string{"This, is", ". a! tes", "t; strin", "g: to", "- split", "_ into", "+ chunks"},
		},
		{
			name:           "No delimiters",
			text:           "Thisisteststring",
			chunkLength:    5,
			rangeLength:    3,
			delimiters:     []rune{},
			expectedResult: []string{"Thisiste", "ststring"},
		},
		{
			name:           "Empty string",
			text:           "",
			chunkLength:    5,
			rangeLength:    3,
			delimiters:     []rune{' '},
			expectedResult: []string{""},
		},
		{
			name: "Long text",
			text: `##Test - Checkout VIP Upsell - Tripadvisor Only - December -  2

* Design reference: [https://www.dropiebox.com/s/dcsvxdsgdfhqg1hnyl26/sample.psd?dl=0](link) (Please note, Dropiebox links may not be accessible)

**Additional Notes:**

* Consider specifying the duration of the test (e.g., entire month of December or specific dates).
* You may want to include a success metric for the test (e.g., percentage increase in VIP signups).`,
			chunkLength: 100,
			rangeLength: 20,
			delimiters:  []rune{' ', '\n'},
			expectedResult: []string{
				`##Test - Checkout VIP Upsell - Tripadvisor Only - December -  2

* Design reference: [https://www.dropiebox.com/s/dcsvxd`,
				`sgdfhqg1hnyl26/sample.psd?dl=0](link) (Please note, Dropiebox links may not be accessible)

**Additional`,
				` Notes:**

* Consider specifying the duration of the test (e.g., entire month of December or specific`,
				` dates).
* You may want to include a success metric for the test (e.g., percentage increase in VIP signups).`,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SplitTextInChunks(tc.text, tc.chunkLength, tc.rangeLength, tc.delimiters)
			if !reflect.DeepEqual(strings.Join(result, ""), tc.text) {
				t.Errorf("The text has changed: %v", tc.text)
			}
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("Expected %v, but got %v", tc.expectedResult, result)
			}

		})
	}
}
