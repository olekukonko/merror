merror (Multiple Error)
==

[![Build Status](https://travis-ci.org/olekukonko/merror.png?branch=master)](https://travis-ci.org/olekukonko/ts) [![Total views](https://sourcegraph.com/api/repos/github.com/olekukonko/merror/counters/views.png)](https://sourcegraph.com/github.com/olekukonko/merror)

Simple Multiple error management in `go`. Some features include :

- Merge Multiple error
- Similar Methods `Append` & `Len`
- Multiple Displays `string` , `line` or `Tab`
- No External dependency


#### To Install

Run `go get github.com/olekukonko/merror` to download and install

#### Example 1

```go
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

	// Append Multiple
	err.Append(errors.New("Error B"),errors.New("Error C"))

	// Print Error
	fmt.Println(err)
	// Output:
	// Error A ; Error B; Error C
}
```


#### Example 2

```go
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
```



[See Documentation](http://godoc.org/github.com/olekukonko/merror)
