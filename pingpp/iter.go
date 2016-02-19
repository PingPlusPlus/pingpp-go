package pingpp

import (
	"net/url"
	"reflect"
)

type Query func(url.Values) ([]interface{}, ListMeta, error)

// list列表数据查询遍历器
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
	return iter
}

//获取当前数据的页
func (it *Iter) getPage() {
	it.values, it.meta, it.err = it.query(it.qs)
	if it.params.End != "" {
		reverse(it.values)
	}
}

//获取下一条数据
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

//获取当前的数据
func (it *Iter) Current() interface{} {
	return it.cur
}

//获取远程数据查询返回错误的信息
func (it *Iter) Err() error {
	return it.err
}

//获取当前遍历器的元属性 has_more/url/object等等
func (it *Iter) Meta() *ListMeta {
	return &it.meta
}

//通过数据的ID属性
func listItemID(x interface{}) string {
	return reflect.ValueOf(x).Elem().FieldByName("ID").String()
}

//分页函数
func reverse(a []interface{}) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
