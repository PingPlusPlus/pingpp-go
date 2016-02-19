package pingpp

import (
	"fmt"
	"net/url"
)

const (
	startafter = "starting_after"
	endbefore  = "ending_before"
)

//数据过滤结构
type Filters struct {
	f []*filter
}

type filter struct {
	Key, Op, Val string
}

//添加一个过滤规则
func (f *Filters) AddFilter(key, op, value string) {
	filter := &filter{Key: key, Op: op, Val: value}
	f.f = append(f.f, filter)
}

//将过滤规则转换成查询参数
func (f *Filters) AppendTo(values *url.Values) {
	for _, v := range f.f {
		if len(v.Op) > 0 {
			values.Add(fmt.Sprintf("%v[%v]", v.Key, v.Op), v.Val)
		} else {
			values.Add(v.Key, v.Val)
		}
	}
}

/* 2016-02-16 当前情况下没有代码调用了该函数
type Params struct {
	Exp   []string
	Meta  map[string]string
	Extra url.Values
}

func (p *Params) Expand(f string) {
	p.Exp = append(p.Exp, f)
}

func (p *Params) AddMeta(key, value string) {
	if p.Meta == nil {
		p.Meta = make(map[string]string)
	}

	p.Meta[key] = value
}

func (p *Params) AddExtra(key, value string) {
	if p.Extra == nil {
		p.Extra = make(url.Values)
	}

	p.Extra.Add(key, value)
}

func (p *Params) AppendTo(body *url.Values) {
	for k, v := range p.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	for _, v := range p.Exp {
		body.Add("expand[]", v)
	}

	for k, vs := range p.Extra {
		for _, v := range vs {
			body.Add(k, v)
		}
	}
}
*/
