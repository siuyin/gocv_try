package main

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestExploreNewMat(t *testing.T) {
	m := gocv.NewMat()
	if d := m.Size(); len(d) != 0 {
		t.Error("dims", d)
	}
}

func TestExploreNewMatFromBytes(t *testing.T) {
	m, err := gocv.NewMatFromBytes(2, 2, gocv.MatTypeCV8U, []byte{1, 0, 0, 1})
	if err != nil {
		t.Error(err)
	}
	if d := m.Size(); d[0] != 2 && d[1] != 2 {
		t.Error("dims", d)
	}
	if e := m.GetUCharAt(0, 0); e != 1 {
		t.Error("0,0:", e)
	}
	if e := m.GetUCharAt(0, 1); e != 0 {
		t.Error("0,1:", e)
	}
}

func TestEncodeFloat(t *testing.T) {
	b := bytesFromFloat32([]float32{1, 0, 0, 1})
	if n := len(b); n != 16 {
		t.Error("len:", n)
	}
}

func TestExploreNewMatFromBytesFloat(t *testing.T) {
	m, err := gocv.NewMatFromBytes(2, 2, gocv.MatTypeCV32F, bytesFromFloat32([]float32{1.0, 0.0, 0.0, 1.0}))
	if err != nil {
		t.Error(err)
	}
	if d := m.Size(); d[0] != 2 && d[1] != 2 {
		t.Error("dims", d)
	}
	if e := m.GetFloatAt(0, 0); e != 1.0 {
		t.Error("0,0:", e)
	}
	if e := m.GetFloatAt(0, 1); e != 0.0 {
		t.Error("0,1:", e)
	}
}
