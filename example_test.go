package merror

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	ErrorSample = []error{errors.New("one"),
		errors.New("two"),
		errors.New("three")}
)

// Basic Usage Example
func ExampleBasic() {
	err := Multi()
	// Append Simple error
	err.Append(errors.New("Error Msg"))

	// Print Error
	fmt.Println(err)
	// Output:
	// Error Msg

}

// Multiple Append Example
func ExampleMultipleAppend() {

	err := Multi()

	// Supports default append
	err.Append(errors.New("Error A"))

	// Support Multiple via append
	err.Append(fmt.Errorf("%s", "Error B"), errors.New("Error C"))

	fmt.Println(err)
	// Output:
	// Error A; Error B; Error C
}

// Check if error exists
func ExampleHasError() {
	err := Multi()

	// Supports default append
	err.Append(errors.New("Error A"))

	// Check If error exist
	if err.HasError() {
		fmt.Println(err)
	}
	// Output:
	// Error A
}

// Using Tab Display
func ExampleTabDisplay() {
	err := Multi()
	buf := new(bytes.Buffer)

	err.Append(ErrorSample...)
	err.Tab(buf)

	fmt.Println(buf.Bytes())
	// Output:
	// [32 51 32 69 114 114 111 114 40 115 41 32 70 111 117 110 100 10 32 32 45 32 111 110 101 10 32 32 45 32 116 119 111 10 32 32 45 32 116 104 114 101 101 10 10]
}

// Using Line Display
func ExampleLineDisplay() {
	err := Multi()
	buf := new(bytes.Buffer)

	err.Append(ErrorSample...)
	err.Line(buf)

	ioutil.WriteFile("out.txt",buf.Bytes(),755)

	fmt.Println(buf.Bytes())
	// Output:
	// [111 110 101 9 10 116 119 111 9 10 116 104 114 101 101 9 10 10]

}
