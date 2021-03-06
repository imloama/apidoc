// SPDX-License-Identifier: MIT

package doc

import (
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/issue9/assert"
)

var _ xml.Unmarshaler = &Callback{}

func TestCallback_UnmarshalXML(t *testing.T) {
	a := assert.New(t)

	obj := &Callback{
		Method:   http.MethodGet,
		Requests: []*Request{{Mimetype: "json", Type: String}},
	}
	str := `<Callback method="GET"><request type="string" mimetype="json"></request></Callback>`

	data, err := xml.Marshal(obj)
	a.NotError(err).Equal(string(data), str)

	obj1 := &Callback{}
	a.NotError(xml.Unmarshal([]byte(str), obj1))
	a.Equal(obj1, obj)

	// 正常
	obj1 = &Callback{}
	str = `<Callback deprecated="1.1.1" method="GET">
		<path path="/users/{id}/orders">
			<param name="id" type="number" summary="用户 ID" />
			<query name="size" type="number" summary="size" />
			<query name="page" type="number" deprecated="0.1.1"><description><![CDATA[page]]></description></query>
		</path>
		<request status="200" mimetype="json" type="object">
			<param name="name" type="string" summary="name" />
			<param name="sex" type="string" summary="sex">
				<enum value="male" summary="male" />
				<enum value="female"><description type="markdown">Female</description></enum>
			</param>
			<param name="age" type="number" summary="age" />
		</request>
	</Callback>`
	a.NotError(xml.Unmarshal([]byte(str), obj1)).
		Equal(obj1.Deprecated, "1.1.1").
		Equal(1, len(obj1.Requests)).
		Equal(obj1.Requests[0].Type, Object).
		NotNil(obj1.Path).
		Equal(obj1.Path.Path, "/users/{id}/orders")

	a.Error(xml.Unmarshal([]byte(str), obj1))

	// 少 method
	str = `<Callback><request type="string" mimetype="json" /></Callback>`
	obj1 = &Callback{}
	a.Error(xml.Unmarshal([]byte(str), obj1))

	// 缺少 request
	obj1 = &Callback{}
	str = `<Callback method="GET" schema="http"></Callback>`
	a.Error(xml.Unmarshal([]byte(str), obj1))

	// 语法错误
	obj1 = &Callback{}
	str = `<Callback name="url" deprecated="x.1.1">text</Callback>`
	a.Error(xml.Unmarshal([]byte(str), obj1))
}
