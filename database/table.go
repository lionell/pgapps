package database

import (
	"bytes"
	"fmt"
	"io"
	"text/tabwriter"
)

type Row []string

type Table struct {
	Header []string `json:"header"`
	Rows   []Row    `json:"rows"`
}

func (t Table) String() string {
	var out bytes.Buffer
	w := tabwriter.NewWriter(&out, 0, 0, 3, ' ', tabwriter.AlignRight)

	t.writeHeader(w)
	for _, r := range t.Rows {
		writeRow(w, r)
	}
	w.Flush()

	return out.String()
}

func (t Table) writeHeader(w io.Writer) {
	var buf bytes.Buffer
	defer buf.WriteTo(w)

	for _, v := range t.Header {
		buf.WriteString(fmt.Sprintf("%v\t", v))
	}
	buf.WriteString("\n")
}

func writeRow(w io.Writer, r Row) {
	var buf bytes.Buffer
	defer buf.WriteTo(w)

	for _, v := range r {
		buf.WriteString(fmt.Sprintf("%v\t", v))
	}
	buf.WriteString("\n")
}
