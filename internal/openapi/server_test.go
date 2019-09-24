// SPDX-License-Identifier: MIT

package openapi

import (
	"testing"

	"github.com/issue9/assert"
)

func TestServer_Sanitize(t *testing.T) {
	a := assert.New(t)

	srv := &Server{}
	a.Error(srv.Sanitize())

	srv.URL = "https://example.com/{tpl1}/{tpl2}/path3"
	a.NotError(srv.Sanitize())

	srv.Variables = map[string]*ServerVariable{
		"tpl1": {Default: "1"},
		"tpl2": {Default: "2", Enum: []string{"1", "2"}},
	}
	a.NotError(srv.Sanitize())

	// variable 不在 URL 中
	srv.Variables = map[string]*ServerVariable{
		"tpl3": {Default: "1"},
	}
	a.Error(srv.Sanitize())

	// variables 存在错误
	srv.Variables = map[string]*ServerVariable{
		"tpl2": {Default: "not-exists", Enum: []string{"1", "2"}},
	}
	a.Error(srv.Sanitize())
}

func TestServerVariable_Sanitize(t *testing.T) {
	a := assert.New(t)

	sv := &ServerVariable{}
	a.Error(sv.Sanitize())

	sv.Enum = []string{"e1", "e2"}
	a.Error(sv.Sanitize())

	sv.Default = "not-in-enum"
	a.Error(sv.Sanitize())

	sv.Default = "e1"
	a.NotError(sv.Sanitize())
}
