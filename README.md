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

#### Example

```go
package main

import (
	"fmt"
	"errors"
	"github.com/olekukonko/merror"
)

func main() {
	var err merror.Errors

	// Append Simple error
	err.Append(errors.New("Error A"))

	// Print Error
	fmt.Println(err)
	// Output:
	// Error A
	// Error B
}
```

[See Documentation](http://godoc.org/github.com/olekukonko/merror)
