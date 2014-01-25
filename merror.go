package merror

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type Errors []error

var (
	mux sync.RWMutex
)

func (e *Errors) Append(err ...error) {
	mux.Lock()
	defer mux.Unlock()
	for _, v := range err {
		*e = append(*e, v)
	}
}

func (e *Errors) Merge(err Errors) {
	mux.Lock()
	defer mux.Unlock()
	for _, v := range err {
		*e = append(*e, v)
	}
}


func (e *Errors) Clear() {
	mux.Lock()
	defer mux.Unlock()
	*e = []error{}
}

func (e *Errors) Len() int {
	mux.Lock()
	defer mux.Unlock()
	return len(*e)
}

func (e Errors) HasError() bool {
	return e.Len() > 0
}

func (e Errors) Display(w io.Writer) Display {
	return Display{e, w}
}

func (e Errors) String() string {
	buf := new(bytes.Buffer)
	l := e.Len()
	for i , v := range e {
		fmt.Fprintf(buf, "%s", v)
		if (i + 1) < l {
			fmt.Fprint(buf, "; ")
		}
	}
	return buf.String()
}
