// Copyright 2016 Mike Scherbakov
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linkedmap

type Element struct {
	next, prev *Element
	list       *LinkedMap
	key        interface{}
}

type LinkedMap struct {
	Map  map[interface{}]interface{}
	root Element
	last *Element
}

func New() *LinkedMap {
	lm := &LinkedMap{}
	lm.Map = make(map[interface{}]interface{})
	lm.root.next = &lm.root
	lm.root.prev = &lm.root
	lm.last = &lm.root
	return lm
}

func (lm *LinkedMap) Add(key interface{}, value interface{}) {
	isUpdate := false
	if lm.Map[key] != nil {
		isUpdate = true // it's actually update, not new added value
	}
	lm.Map[key] = value

	if isUpdate {
		return
	}

	// If we add new k,v pairs only
	e := &Element{nil, nil, lm, key}

	lm.last.next = e
	e.prev = lm.last
	lm.last = e
}

func (lm *LinkedMap) Get(key interface{}) interface{} {
	return lm.Map[key]
}

func (e *Element) Key() interface{} {
	return e.key
}

func (e *Element) Value() interface{} {
	return e.list.Get(e.key)
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	if e.prev != &e.list.root {
		return e.prev
	}
	return nil
}

func (lm *LinkedMap) First() *Element {
	return lm.root.next
}

func (lm *LinkedMap) Last() *Element {
	return lm.last
}
