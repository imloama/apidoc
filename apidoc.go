// SPDX-License-Identifier: MIT

// Package apidoc RESTful API 文档生成工具
//
// 可以从代码文件的注释中提取文档内容，生成 API 文档，
// 支持大部分的主流的编程语言。
package apidoc

import (
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/language"

	"github.com/caixw/apidoc/v5/input"
	"github.com/caixw/apidoc/v5/internal/docs"
	"github.com/caixw/apidoc/v5/internal/locale"
	"github.com/caixw/apidoc/v5/internal/vars"
	"github.com/caixw/apidoc/v5/message"
	"github.com/caixw/apidoc/v5/output"
)

// Init 初始化包
//
// 如果传递了 language.Und，则采用系统当前的本地化信息。
// 如果获取系统的本地化信息依然失败，则会失放 zh-Hans 作为默认值。
func Init(tag language.Tag) error {
	return locale.Init(tag)
}

// Version 当前程序的版本号
//
// 为一个正常的 semver(https://semver.org/lang/zh-CN/) 格式字符串。
func Version() string {
	return vars.Version()
}

// Do 解析文档并输出文档内容
//
// 如果需要控制详细的操作步骤，可以自行调用 input 和 output 的相关函数实现。
//
// 如果是文档语法错误，则相关的错误信息会反馈给 h，由 h 处理错误信息；
// 如果是配置项（o 和 i）有问题，则以 *message.SyntaxError 类型返回错误信息。
func Do(h *message.Handler, o *output.Options, i ...*input.Options) error {
	doc, err := input.Parse(h, i...)
	if err != nil {
		return err
	}

	return output.Render(doc, o)
}

// Site 将 dir 作为静态文件服务内容
//
// 默认页为 index.xml，同时会过滤 CNAME，
// 如果将 dir 指同 docs 目录，相当于本地版本的 https://apidoc.tools
//
// 用户可以通过诸如：
//  http.Handle("/apidoc", apidoc.Site("./docs"))
// 的代码搭建一个简易的 https://apidoc.tools 网站。
func Site(dir string) http.Handler {
	return docs.Handler(dir)
}

// Handle 返回显示文档内容的中间件
//
// p 指定了 apidoc.xml 实际的文件路径；
// contentType 表示 p 的 mimetype 类型，如果为空，则会采用 "application/xml";
// l 表示出错时，错误内容的发送通道，如果为 nil，表示不输出错误信息；
func Handle(p, contentType string, l *log.Logger) http.Handler {
	if contentType == "" {
		contentType = "application/xml"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		printErr := func(err error) {
			if l != nil {
				l.Println(err)
			}
		}

		data, err := ioutil.ReadFile(p)
		if err != nil {
			printErr(err)
			errStatus(w, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", contentType)
		if _, err = w.Write(data); err != nil {
			printErr(err) // 此时 writeHeader 已经发出，再输出状态码无意义
		}
	})
}

func errStatus(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
