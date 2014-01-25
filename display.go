package merror

import (
	"fmt"
	"io"
)


type Display struct {
	e Errors
	w io.Writer
}

func (d Display) Tab() {
	fmt.Fprintf(d.w, "\n%-1s%d Error(s) Found \n", " " , d.e.Len())
	for _ , v := range d.e {
		fmt.Fprintf(d.w, "%-2s%-2s%s\t\n", " ", "-", v)
	}
	fmt.Fprintln(d.w)
}

func (d Display) Line() {
	for _, v := range d.e {
		fmt.Fprintf(d.w, "%s\t\n", v)
	}
	fmt.Fprintln(d.w)
}
