package merror

import (
	"errors"
	"testing"
)


func TestEmpty(t *testing.T) {
	var err Errors
	if err != nil {
		t.Error("An Errors should be Nill")
	}
}

func TestLength(t *testing.T) {
	var err Errors
	err.Append(ErrorSample...)

	if len(err) != 3 {
		t.Error("Error Lenght does not match")
	}

	if len(err) != err.Len() {
		t.Error("len & Len should be same value")
	}

	if !err.HasError() {
		t.Error("HasError Failed to detect error")
	}

}

func TestClear(t *testing.T) {
	var err Errors
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

func TestAppend(t *testing.T) {
	var err Errors
	err.Append(ErrorSample...)
	err = append(err, ErrorSample...)
	if err.Len() != 6 {
		t.Error("Multiple Append should work")
	}
}

func TestRace(t *testing.T) {
	var err Errors

	for i := 0; i < 10; i++ {
		go func() {
			err.Append(ErrorSample...)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			err.Len()
		}()
	}

}

func TestString(t *testing.T) {
	var err Errors
	err.Append(ErrorSample...)
	expect := "one; two; three"
	if expect != err.String() {
		t.Error("Output String not consitent")
	}
}



func TestMerge(t *testing.T) {
	var err1 , err2  Errors

	err1.Append(ErrorSample...)
	err2.Append(errors.New("four"))

	expect := "one; two; three; four"

	err1.Merge(err2)

	if expect != err1.String() {
		t.Error("Merge String not consitent")
	}
}


