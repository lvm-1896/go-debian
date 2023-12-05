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

package control_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"pault.ag/go/debian/control"
)

func TestSourcesMarshal(t *testing.T) {
	testStruct := TestMarshalStruct{Foo: `Hello
This
Is

A Test`}
	writer := bytes.Buffer{}

	err := control.Marshal(&writer, testStruct)
	isok(t, err)

	assert(t, writer.String() == `Foo: Hello
 This
 Is
 .
 A Test
`)
}

/*
 *
 */

func TestSourcesParse(t *testing.T) {
	// Test Control {{{
	reader := bufio.NewReader(strings.NewReader(`# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.

## Ubuntu distribution repository
##
## The following settings can be tweaked to configure which packages to use from Ubuntu.
## Mirror your choices (except for URIs and Suites) in the security section below to
## ensure timely security updates.
## 
## Types: Append deb-src to enable the fetching of source package.
## URIs: A URL to the repository (you may add multiple URLs)
## Suites: The following additional suites can be configured
##   <name>-updates   - Major bug fix updates produced after the final release of the
##                      distribution.
##   <name>-backports - software from this repository may not have been tested as
##                      extensively as that contained in the main release, although it includes
##                      newer versions of some applications which may provide useful features.
##                      Also, please note that software in backports WILL NOT receive any review
##                      or updates from the Ubuntu security team.
## Components: Aside from main, the following components can be added to the list
##   restricted  - Software that may not be under a free license, or protected by patents.
##   universe    - Community maintained packages. Software from this repository is 
##                 ENTIRELY UNSUPPORTED by the Ubuntu team. Also, please note
##                 that software in universe WILL NOT receive any
##                 review or updates from the Ubuntu security team.
##   multiverse  - Community maintained of restricted. Software from this repository is
##                 ENTIRELY UNSUPPORTED by the Ubuntu team, and may not be under a free 
##                 licence. Please satisfy yourself as to your rights to use the software.
##                 Also, please note that software in multiverse WILL NOT receive any 
##                 review or updates from the Ubuntu security team.
##
## See the sources.list(5) manual page for further settings.
Types: deb
URIs: http://archive.ubuntu.com/ubuntu
Suites: lunar lunar-updates
Components: main universe
Signed-By: /usr/share/keyrings/ubuntu-archive-keyring.gpg

## Ubuntu security updates. Aside from URIs and Suites,
## this should mirror your choices in the previous section.
Types: deb
URIs: http://security.ubuntu.com/ubuntu
Suites: lunar-security
Components: main universe
Signed-By: /usr/share/keyrings/ubuntu-archive-keyring.gpg

Types: deb
URIs: http://archive.ubuntu.com/ubuntu
Suites: lunar-proposed
Components: main universe
Signed-By: /usr/share/keyrings/ubuntu-archive-keyring.gpg
`))
	// }}}
	c, err := control.ParseDebSources(reader, "")
	//fmt.Println(err)
	isok(t, err)
	assert(t, c != nil)
	//  fmt.Println(c)
	// fmt.Println(c.DebSources[0].Types)
	fmt.Println(c.DebSources[0].Suites)
	fmt.Println(c.DebSources[1].Suites)
	fmt.Println(c.DebSources[2].Suites)
	fmt.Println(c.DebSources[1].Order)
	writer := bytes.Buffer{}
	err = control.Marshal(&writer, *c)
	isok(t, err)
	fmt.Println(writer.Bytes())

	assert(t, len(c.DebSources) == 2)

	// assert(t, c.Source.Maintainer == "Paul Tagliamonte <paultag@ubuntu.com>")
	// assert(t, c.Source.Source == "fbautostart")

	//depends := c.DebSource

	// assert(t, len(c.Source.Maintainers()) == 1)
	assert(t, len(c.DebSources[0].Types) == 1)
}

// vim: foldmethod=marker
