// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package scanner

import (
	"testing"

	"github.com/issue9/assert"
)

func TestLangs(t *testing.T) {
	a := assert.New(t)

	ls := Langs()
	a.Equal(len(ls), len(langs))
}

func TestExtsIndex(t *testing.T) {
	a := assert.New(t)

	a.Equal(extsIndex[".cpp"], "cpp")
	a.Equal(extsIndex[".php"], "php")
	a.Equal(extsIndex[".go"], "go")
}
