package base64_test

import (
	"bytes"
	stdbase64 "encoding/base64"
	"fmt"
	"testing"

	"github.com/ymgyt/codeconv/internal/base64"
)

func TestConverter_Encode(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{[]byte(``), ""},
		{[]byte(`hello gopher`), "aGVsbG8gZ29waGVy"},
	}

	c := buildConverter()

	for i, tc := range tests {
		t.Run(fmt.Sprintf("case: %d", i), func(t *testing.T) {
			var w bytes.Buffer
			err := c.Encode(&w, bytes.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			got, want := w.String(), tc.want
			if got != want {
				t.Errorf("Encode() = %s; want %s", got, want)
			}
		})
	}
}

func TestConverter_Decode(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{[]byte("aGVsbG8gZ29waGVy"), "hello gopher"},
	}

	c := buildConverter()

	for i, tc := range tests {
		t.Run(fmt.Sprintf("case: %d", i), func(t *testing.T) {
			var w bytes.Buffer
			err := c.Decode(&w, bytes.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			got, want := w.String(), tc.want
			if got != want {
				t.Errorf("Decode() = %s; want %s", got, want)
			}
		})
	}
}

func buildConverter() *base64.Converter {
	return &base64.Converter{
		Opt: &base64.Options{Encoding: stdbase64.StdEncoding},
	}
}
