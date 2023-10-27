package lotrsdk

import "strings"

type Filter struct {
	params map[string]string
}

func NewFilter() *Filter {
	return &Filter{
		params: make(map[string]string),
	}
}

func (f *Filter) Add(key, value string) {
	f.params[key+"="] = value
}

func (f *Filter) Remove(key string) {
	delete(f.params, key)
}

func (f *Filter) AddNotEqual(key, value string) {
	f.params[key+"!="] = value
}

func (f *Filter) AddGreaterThan(key, value string) {
	f.params[key+">"] = value
}

func (f *Filter) AddLessThan(key, value string) {
	f.params[key+"<"] = value
}

func (f *Filter) AddGreaterThanOrEqualTo(key, value string) {
	f.params[key+">="] = value
}

func (f *Filter) AddLessThanOrEqualTo(key, value string) {
	f.params[key+"<="] = value
}

func (f *Filter) AddMatchRegex(key, pattern string) {
	f.params[key+"="] = "/" + pattern + "/"
}

func (f *Filter) AddNotMatchRegex(key, pattern string) {
	f.params[key+"!="] = "/" + pattern + "/"
}

func (f *Filter) AddFieldExist(field string) {
	f.params[field+"?"] = ""
}

func (f *Filter) AddFieldNotExist(field string) {
	f.params[field+"?!"] = ""
}

func (f *Filter) Encode() string {
	var parts []string
	for k, v := range f.params {
		parts = append(parts, k+v)
	}
	return strings.Join(parts, "&")
}
