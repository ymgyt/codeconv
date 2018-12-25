package base64

import (
	"encoding/base64"
	"io"
	"io/ioutil"
)

// Options -
type Options struct {
	Encoding *base64.Encoding
}

// Converter convert to/from base64 format.
type Converter struct {
	Opt *Options
}

// Encode -
func (c *Converter) Encode(w io.Writer, r io.Reader) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	enc := base64.NewEncoder(c.Opt.Encoding, w)
	defer func() {
		if err == nil {
			err = enc.Close()
		}
	}()
	_, err = enc.Write(b)
	return
}

// Decode -
func (c *Converter) Decode(w io.Writer, r io.Reader) (err error) {
	dec := base64.NewDecoder(c.Opt.Encoding, r)
	_, err = io.Copy(w, dec)
	return err
}
