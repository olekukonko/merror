// Copyright 2014 Oleku Konko All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This module is a Terminal  API for the Go Programming Language.
// The protocols were written in pure Go and works on windows and unix systems
package merror

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

// Error Type
type Errors struct {
	lock  *sync.RWMutex
	items []error
}


// Initiate Error Instance
func Multi() *Errors {
	return &Errors{new(sync.RWMutex), []error{}}
}

// Append new error
func (e *Errors) Append(err ...error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	for _, v := range err {
		e.items = append(e.items, v)
	}
}

// Merge Errors
func (e *Errors) Merge(err *Errors) {
	e.Append(err.List()...)
}

// List all error items
func (e Errors) List() []error {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.items
}

// Clear all error items
func (e *Errors) Clear() {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.items = []error{}
}

// Total lent of error in list
func (e Errors) Len() int {
	e.lock.Lock()
	defer e.lock.Unlock()
	return len(e.items)
}

// Check if it has errors
func (e Errors) HasError() bool {
	return e.Len() > 0
}

// Return Stringer interface
func (e Errors) String() string {
	buf := new(bytes.Buffer)
	l := e.Len()
	for i, v := range e.List() {
		fmt.Fprintf(buf, "%s", v)
		if (i + 1) < l {
			fmt.Fprint(buf, "; ")
		}
	}
	return buf.String()
}

// Display errors in tabulated format
func (e Errors) Tab(w io.Writer) {
	fmt.Fprintf(w, "%-1s%d Error(s) Found\n", " ", e.Len())
	for _, v := range e.List() {
		fmt.Fprintf(w, "%-2s%-2s%s\n", " ", "-", v)
	}
	fmt.Fprintln(w)
}

// Display errors in Line format
func (e Errors) Line(w io.Writer) {
	for _, v := range e.List() {
		fmt.Fprintf(w, "%s\t\n", v)
	}
	fmt.Fprintln(w)
}
