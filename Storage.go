package repository

import (
	"fmt"
	"sync"
)

func NewStorage(name string) *Storage { // {{{
	return &Storage{
		name:      name,
		data:      make(Data, 0),
		indexes:   make(Indexes, 0),
		callbacks: make(Callbacks, 0),
	}
} // }}}

type Storage struct {
	sync.RWMutex

	name      string
	data      Data
	indexes   Indexes
	callbacks Callbacks
}

// [Protected]

func (this *Storage) updateIndexes(data interface{}, i int) { // {{{
	this.RLock()
	defer this.RUnlock()
	for name, index := range this.indexes {
		callback := this.callbacks[name]
		index.Set(callback(data, i), i)
	}
} // }}}

// [Public]

func (this *Storage) Name() string { // {{{
	return this.name
} // }}}

func (this *Storage) HasIndex(name string) bool { // {{{
	_, ok := this.indexes[name]
	return ok
} // }}}

func (this *Storage) Index(name string) (IIndex, error) { // {{{
	if !this.HasIndex(name) {
		return nil, fmt.Errorf("Index '%s' is not exist", name)
	}

	return this.indexes[name], nil
} // }}}

func (this *Storage) CreateIndex(name string, callback Callback) error { // {{{
	return this.CreateOneToOneIndex(name, callback)
} // }}}

func (this *Storage) CreateOneToOneIndex(name string, callback Callback) error { // {{{
	if this.HasIndex(name) {
		return fmt.Errorf("Index '%s' is already exist", name)
	}
	this.indexes[name] = NewIndex()
	this.callbacks[name] = callback

	return nil
} // }}}

func (this *Storage) CreateOneToManyIndex(name string, callback Callback) error { // {{{
	// @TODO:
	return nil
} // }}}

func (this *Storage) CreateManyToOneIndex(name string, callback Callback) error { // {{{
	// @TODO:
	return nil
} // }}}

func (this *Storage) CreateManyToManyIndex(name string, callback Callback) error { // {{{
	// @TODO:
	return nil
} // }}}

func (this *Storage) Data() Data { // {{{
	return this.data
} // }}}

func (this *Storage) Add(data interface{}) { // {{{
	this.Lock()
	this.data = append(this.data, data)
	i := len(this.data)
	this.Unlock()

	this.updateIndexes(data, i)
} // }}}

func (this *Storage) Get(indexName string, indexValue interface{}) (interface{}, error) { // {{{
	if !this.HasIndex(indexName) {
		return nil, fmt.Errorf("Index '%s' is not exist", indexName)
	}

	if i, ok := this.indexes[indexName].Get(indexValue); ok {
		return this.data[i-1], nil
	}

	return nil, fmt.Errorf("Cannot find '%v' in index '%s'", indexValue, indexName)
} // }}}

func (this *Storage) MatchFirst(indexName string, pattern string) (interface{}, error) { // {{{
	if !this.HasIndex(indexName) {
		return nil, fmt.Errorf("Index '%s' is not exist", indexName)
	}

	if i, ok := this.indexes[indexName].MatchFirst(pattern); ok {
		return this.data[i-1], nil
	}

	return nil, fmt.Errorf("Not matched '%s' in index '%s'", pattern, indexName)
} // }}}
