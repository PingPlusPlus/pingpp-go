package pingpp

import (
	"fmt"
	"net/url"
	"reflect"
)

type Query func(url.Values) ([]interface{}, ListMeta, error)

type Iter struct {
	query  Query
	qs     url.Values
	values []interface{}
	meta   ListMeta
	params ListParams
	err    error
	cur    interface{}
}

func GetIter(params *ListParams, qs *url.Values, query Query) *Iter {
	iter := &Iter{}
	iter.query = query

	p := params

	if p == nil {
		p = &ListParams{}
	}
	iter.params = *p
	q := qs
	if q == nil {
		q = &url.Values{}
	}
	iter.qs = *q

	iter.getPage()
	fmt.Println(iter)
	return iter
}

func (it *Iter) getPage() {
	it.values, it.meta, it.err = it.query(it.qs)
	if it.params.End != "" {
		reverse(it.values)
	}
}

func (it *Iter) Next() bool {
	if len(it.values) == 0 && it.meta.More && !it.params.Single {
		if it.params.End != "" {
			it.params.End = listItemID(it.cur)
			it.qs.Set(endbefore, it.params.End)
		} else {
			it.params.Start = listItemID(it.cur)
			it.qs.Set(startafter, it.params.Start)
		}
		it.getPage()
	}
	if len(it.values) == 0 {
		return false
	}
	it.cur = it.values[0]
	it.values = it.values[1:]
	return true
}

func (it *Iter) Current() interface{} {
	return it.cur
}

func (it *Iter) Err() error {
	return it.err
}

func (it *Iter) Meta() *ListMeta {
	return &it.meta
}

func listItemID(x interface{}) string {
	return reflect.ValueOf(x).Elem().FieldByName("ID").String()
}

func reverse(a []interface{}) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
