// Copyright 2014 Oleku Konko All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This module is a Terminal  API for the Go Programming Language.
// The protocols were written in pure Go and works on windows and unix systems

/**

Simple Multiple error management in `go`. Some features include :

- Merge Multiple error
- Similar Methods `Append` & `Len`
- Multiple Displays `string` , `line` or `Tab`
- No External dependency

To Install

Run `go get github.com/olekukonko/merror` to download and install

Example 1

	package main

	import (
		"fmt"
		"errors"
		"github.com/olekukonko/merror"
	)

	func main() {

		// new Multiple Error
		err := merror.Multi()

		// Append Simple error
		err.Append(errors.New("Error A"))

		// Use default append
		err = append(err,errors.New("Error B"))

		// Print Error
		fmt.Println(err)
		// Output:
		// Error A; Error B
	}


Example 2


	package main

	import (
		"fmt"
		"errors"
		"github.com/olekukonko/merror"
		"os"
	)

	func main() {
		// new Multiple Error
		err := merror.Multi()

		// Append Simple error
		err.Append(errors.New("Error A"))

		// Append Multiple
		err.Append(errors.New("Error B"),errors.New("Error C"))

		// Send to output command line or http
		err.Tab(os.Stdout)
		// Output:
		// 3 Error(s) Found
		//     - Error A
		//     - Error B
		//     - Error C
	}


**/
package merror
