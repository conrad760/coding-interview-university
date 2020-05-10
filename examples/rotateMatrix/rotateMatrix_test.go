package main

import (
	"reflect"
	"testing"
)

func TestMatrixRotates(t *testing.T) {

	actual := randMatrix(3)
	expected := [][]int{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}}
	clockwise(actual)
	printMatrix(actual)
	printMatrix(expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Error("oh no!")
	}

}
