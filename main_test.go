package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	caseTable:=[]struct{
		test []int
		expected []int
	}{
		{
			test:     []int{},
			expected:   []int{},
		},
		{
			test:     []int{0,0,0,0},
			expected:   []int{0,0,0,0},
		},
		{
			test:     []int{5,3,412,32,-1,2,-3,-2},
			expected:   []int{-3,-2,-1,2,3,5,32,412},
		},
	}

	for _, testCase := range caseTable {
		t.Logf("%v", testCase.test)
		result := MergeSort(testCase.test)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result: %v, expected %v", result, testCase.expected)
		}
	}
}
