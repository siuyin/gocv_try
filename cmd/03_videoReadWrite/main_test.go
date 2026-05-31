package main

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestMatGetDimensions(t *testing.T) {
	m := gocv.NewMatWithSize(4, 3, gocv.MatTypeCV8U)
	if &m == nil {
		t.Error("mat should not be nil")
	}
	if r := m.Rows(); r != 4 {
		t.Errorf("rows: %v should be 4", r)
	}
	if c := m.Cols(); c != 3 {
		t.Errorf("columns: %v should be 3", c)
	}
}
