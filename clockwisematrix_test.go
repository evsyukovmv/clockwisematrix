package clockwisematrix

import (
	"reflect"
	"testing"
)

type testInputOutput struct {
	input  [][]int
	output []int
}

var testData = [...]testInputOutput{
	{
		input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		input: [][]int{
			{1, 2, 3, 1},
			{4, 5, 6, 4},
			{7, 8, 9, 7},
			{7, 8, 9, 7},
		},
		output: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
	},
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{7, 8, 9, 10, 11},
			{13, 14, 15, 16, 17},
			{19, 20, 21, 22, 23},
			{25, 26, 27, 28, 29},
		},
		output: []int{1, 2, 3, 4, 5, 11, 17, 23, 29, 28, 27, 26, 25, 19, 13, 7, 8, 9, 10, 16, 22, 21, 20, 14, 15},
	},
	{
		input: [][]int{
			{1, 2, 3, 4, 5, 6},
			{7, 8, 9, 10, 11, 12},
			{13, 14, 15, 16, 17, 18},
			{19, 20, 21, 22, 23, 24},
			{25, 26, 27, 28, 29, 30},
			{31, 32, 33, 34, 35, 36},
		},
		output: []int{1, 2, 3, 4, 5, 6, 12, 18, 24, 30, 36, 35, 34, 33, 32, 31, 25, 19, 13, 7, 8, 9, 10, 11, 17, 23, 29, 28, 27, 26, 20, 14, 15, 16, 22, 21},
	},
	{
		input:  [][]int{},
		output: []int{},
	},
}

func TestFlatten(t *testing.T) {
	for _, data := range testData {
		result, _ := Flatten(data.input)
		if !reflect.DeepEqual(result, data.output) {
			t.Errorf("test failed - results not match\nGot:\n%v\nExpected:\n%v", result, data.output)
			break
		}
	}
}

func TestFlattenErrors(t *testing.T) {
	_, err := Flatten([][]int{{1, 2, 3}, {4, 5, 6}})
	errorString := "not valid, square matrix required"
	if err.Error() != errorString {
		t.Errorf("test failed - errors not match\nGot:\n%v\nExpected:\n%v", err.Error(), errorString)
	}
}
