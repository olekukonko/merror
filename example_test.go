package merror

import (
	"errors"
	"fmt"
)

var (
	ErrorSample = []error{errors.New("one"),
	errors.New("two"),
	errors.New("three")}
)

// Basic Usage Example
func ExampleBasic() {
	var err Errors
	// Append Simple error
	err.Append(errors.New("Error Msg"))

	// Print Error
	fmt.Println(err)
	// Output:
	// Error Msg

}

// Multiple Append Example
func ExampleMultipleAppend() {

	err := Errors{}

	// Supports default append
	err = append(err, errors.New("Error A"))

	// Support Multiple via append
	err.Append(fmt.Errorf("%s", "Error B"), errors.New("Error C"))

	fmt.Println(err)
	// Output:
	// Error Error A; Error B; Error C
}

// Check if error exists
func ExampleHasError() {
	var err Errors

	// Supports default append
	err.Append(errors.New("Error A"))

	// Check If error exist
	if err.HasError() {
		fmt.Println(err)
	}
	// Output:
	// Error Error A; Error B; Error C
}

// Using Error in Multiple Goroutines
func ExampleRace() {
	var err  Errors
	for i := 0 ; i < 3 ; i ++ {
		go func(e Errors) {
			e.Append(fmt.Errorf("Error %d", i))
		}(err)
	}

	fmt.Println(err)

	// Output:
	// Error Error 1; Error 2; Error 3
}

// Using Error Display
func ExampleDisplay() {
	var err  Errors
	err.Append(ErrorSample...)
	fmt.Println(err)
	// Output:
	//	3  Error(s) Found
	//	- one
	//	- two
	//	- three
}








