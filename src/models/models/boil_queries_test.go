// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"regexp"

	"github.com/volatiletech/sqlboiler/boil"
)

var dbNameRand *rand.Rand

func MustTx(transactor boil.Transactor, err error) boil.Transactor {
	if err != nil {
		panic(fmt.Sprintf("Cannot create a transactor: %s", err))
	}
	return transactor
}

var rgxPGFkey = regexp.MustCompile(`(?m)^ALTER TABLE ONLY .*\n\s+ADD CONSTRAINT .*? FOREIGN KEY .*?;\n`)
var rgxMySQLkey = regexp.MustCompile(`(?m)((,\n)?\s+CONSTRAINT.*?FOREIGN KEY.*?\n)+`)
var rgxMSSQLkey = regexp.MustCompile(`(?m)^ALTER TABLE .*ADD\s+CONSTRAINT .* FOREIGN KEY.*?.*\n?REFERENCES.*`)

func newFKeyDestroyer(regex *regexp.Regexp, reader io.Reader) io.Reader {
	return &fKeyDestroyer{
		reader: reader,
		rgx:    regex,
	}
}

type fKeyDestroyer struct {
	reader io.Reader
	buf    *bytes.Buffer
	rgx    *regexp.Regexp
}

func (f *fKeyDestroyer) Read(b []byte) (int, error) {
	if f.buf == nil {
		all, err := ioutil.ReadAll(f.reader)
		if err != nil {
			return 0, err
		}

		f.buf = bytes.NewBuffer(f.rgx.ReplaceAll(all, []byte{}))
	}

	return f.buf.Read(b)
}
