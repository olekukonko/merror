// Copyright 2014 Oleku Konko All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This module is a Terminal  API for the Go Programming Language.
// The protocols were written in pure Go and works on windows and unix systems
package merror

import (
	"bytes"
	"errors"
	//"fmt"
	"testing"
)

// Test error item length
func TestLength(t *testing.T) {
	err := Multi()
	err.Append(ErrorSample...)

	//fmt.Println("Total", err.List())

	if len(err.List()) != 3 {
		t.Error("Error Lenght does not match")
	}

	if len(err.List()) != err.Len() {
		t.Error("List Lenght & Len should be same value")
	}

	if !err.HasError() {
		t.Error("HasError Failed to detect error")
	}

}

// Test ability to clear error items
func TestClear(t *testing.T) {
	err := Multi()
	err.Append(ErrorSample...)
	err.Clear()
	if err.HasError() {
		t.Error("There should be no errors after clear")
	}
	err.Append(ErrorSample[0])

	if !err.HasError() {
		t.Error("There should be error after append")
	}
}

// Test Multiple Append
func TestAppend(t *testing.T) {
	err := Multi()
	err.Append(ErrorSample...)
	err.Append(ErrorSample...)

	if err.Len() != 6 {
		t.Error("Multiple Append should work")
	}
}

// Test String Output
func TestString(t *testing.T) {
	err := Multi()
	err.Append(ErrorSample...)
	expect := "one; two; three"
	if expect != err.String() {
		t.Error("Output String not consitent")
	}
}

// Test error merging
func TestMerge(t *testing.T) {

	err1, err2 := Multi(), Multi()

	err1.Append(ErrorSample...)
	err2.Append(errors.New("four"))

	expect := "one; two; three; four"

	err1.Merge(err2)

	if expect != err1.String() {
		t.Error("Merge String not consitent")
	}
}

// Test Tab Output
func TestTab(t *testing.T) {
	err := Multi()
	buf := new(bytes.Buffer)

	err.Append(ErrorSample...)
	err.Tab(buf)

	if buf.String() == "" {
		t.Error("Tab Result Should not be empty")
	}
}

// Test Line Output
func TestLine(t *testing.T) {
	err := Multi()
	buf := new(bytes.Buffer)

	err.Append(ErrorSample...)
	err.Line(buf)

	if buf.String() == "" {
		t.Error("Line Result Shwould not be empty")
	}
}
