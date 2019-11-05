// SPDX-License-Identifier: MIT

package output

import (
	"testing"

	"github.com/issue9/assert"
)

func TestStylesheet(t *testing.T) {
	a := assert.New(t)
	a.NotEmpty(stylesheetURL)
}

func TestOptions_contains(t *testing.T) {
	a := assert.New(t)

	o := &Options{}
	a.True(o.contains("tag"))
	a.True(o.contains(""))

	o.Tags = []string{"t1", "t2"}
	a.True(o.contains("t1"))
	a.False(o.contains("not-exists"))
	a.False(o.contains(""))
}

func TestOptions_sanitize(t *testing.T) {
	a := assert.New(t)

	var o *Options
	a.Error(o.sanitize(false))

	o = &Options{}
	a.Error(o.sanitize(false))

	a.NotError(o.sanitize(true))

	o.Path = "./testdir/apidoc.json"
	a.NotError(o.sanitize(false))
	a.Equal(o.Style, stylesheetURL).
		Equal(2, len(o.procInst)).
		Contains(o.procInst[1], stylesheetURL)
}
