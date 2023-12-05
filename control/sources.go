/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@debian.org>, 2015
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

package control // import "pault.ag/go/debian/control"

import (
	"bufio"
	// "fmt"
	"os"
	"path/filepath"
)

// Encapsulation for a debian/sources file, which is a series of RFC2822-like
// blocks, a series of Source paragraphs.
type DebSources struct {
	Filename string

	DebSources []DebSourcesParagraph
}

// Encapsulation for a debian/control Source entry. This contains information
// that will wind up in the .dsc and friends. Really quite fun!
type DebSourcesParagraph struct {
	Paragraph

	Types      []string `delim:" "`
	URIs       string
	Suites     []string `delim:" "`
	Components []string `delim:" "`
	SignedBy   string   `control:"Signed-By"`
}

// Given a path on the filesystem, Parse the file off the disk and return
// a pointer to a brand new Control struct, unless error is set to a value
// other than nil.
func ParseDebSourcesFile(path string) (ret *DebSources, err error) {
	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ret, err = ParseDebSources(bufio.NewReader(f), path)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Given a bufio.Reader, consume the Reader, and return a Control object
// for use.
func ParseDebSources(reader *bufio.Reader, path string) (*DebSources, error) {
	ret := DebSources{
		Filename:   path,
		DebSources: []DebSourcesParagraph{},
	}

	if err := Unmarshal(&ret.DebSources, reader); err != nil {
		return nil, err
	}

	return &ret, nil
}

// vim: foldmethod=marker
