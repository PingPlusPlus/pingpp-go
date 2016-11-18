package pingpp

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	startafter = "starting_after"
	endbefore  = "ending_before"
)

// 列表查询请求参数
type ListParams struct {
	Start, End string
	Limit      int
	Filters    Filters

	Single bool
}

type PagingParams struct {
	Filters Filters
}

func (p *ListParams) AppendTo(body *url.Values) {
	if len(p.Filters.f) > 0 {
		p.Filters.AppendTo(body)
	}

	if len(p.Start) > 0 {
		body.Add(startafter, p.Start)
	}

	if len(p.End) > 0 {
		body.Add(endbefore, p.End)
	}

	if p.Limit > 0 {
		if p.Limit > 100 {
			p.Limit = 100
		}

		body.Add("limit", strconv.Itoa(p.Limit))
	}
}

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
