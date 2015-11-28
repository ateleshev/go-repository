package repository

import (
	"regexp"
	"sync"
)

func NewIndex() *Index { // {{{
	return &Index{
		data: make(IndexData, 0),
	}
} // }}}

type Index struct {
	IIndex
	sync.RWMutex

	data IndexData
}

func (this *Index) Data() IndexData { // {{{
	return this.data
} // }}}

func (this *Index) Values() Data { // {{{
	values := make(Data, 0)
	for value, _ := range this.data {
		values = append(values, value)
	}

	return values
} // }}}

func (this *Index) Set(key interface{}, value int) { // {{{
	this.Lock()
	defer this.Unlock()

	this.data[key] = value
} // }}}

func (this *Index) Has(key interface{}) bool { // {{{
	_, ok := this.data[key]
	return ok
} // }}}

func (this *Index) Get(key interface{}) (int, bool) { // {{{
	if i, ok := this.data[key]; ok {
		return i, true
	}

	return 0, false
} // }}}

func (this *Index) MatchFirst(pattern string) (int, bool) { // {{{
	for key, value := range this.data {
		if matched, _ := regexp.MatchString(pattern, key.(string)); matched {
			return value, true
		}
	}

	return 0, false
} // }}}
