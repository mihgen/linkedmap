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

import "testing"

func TestElementKeyUpdated(t *testing.T) {
	lm := New()
	key := "key"
	lm.Add(key, "value")

	if lm.last.key != key {
		t.Errorf("Expected element's key %v, got %v", key, lm.last.key)
	}
	got := lm.last.Key()
	if got != key {
		t.Errorf("Expected Key()=%v, got %v", key, got)
	}
}

func TestEmpty(t *testing.T) {
	lm := New()
	if lm.Last() != nil {
		t.Error("Last has to be nil, but it is", lm.Last())
	}
	if lm.First() != nil {
		t.Error("First has to be nil, but it is", lm.First())
	}
}

func TestNilKey(t *testing.T) {
	lm := New()
	lm.Add(nil, "value")
	if lm.Get(nil) != "value" {
		t.Error("Get(nil) must work and return value")
	}
	if lm.last.Key() != nil {
		t.Error("nil key must be stored and retrieved as nil")
	}
}

func TestFirstLast(t *testing.T) {
	lm := New()
	lm.Add(1, "v")
	e := lm.last
	if e != lm.First() {
		t.Error("Wrong element returned for only one stored")
	}
	if e != lm.Last() {
		t.Error("Wrong element returned for only one stored")
	}
}

func TestOneElementRefersToNil(t *testing.T) {
	lm := New()
	lm.Add(-4, true)
	e := lm.last
	if e.next != nil {
		t.Error("e.next must not be defined for one element")
	}
	if e.prev != nil {
		t.Error("e.prev must not be defined for one element")
	}
}

func TestPrevElement(t *testing.T) {
	lm := New()
	lm.Add(0, 0)
	e := lm.last
	if e.Prev() != nil {
		t.Errorf("Prev must not be defined for single element, got %v", e.Prev())
	}
}

func TestNextElement(t *testing.T) {
	lm := New()
	lm.Add(0, 0)
	e := lm.last
	if e.Next() != nil {
		t.Errorf("Next must not be defined for last element, got %v", e.Next())
	}
}

func TestElementValue(t *testing.T) {
	lm := New()
	expected := "value"
	lm.Add("k", expected)

	v := lm.Last().Value()
	if v != expected {
		t.Errorf("Element value expected %v, got %v", expected, v)
	}
	v = lm.Get("k")
	if v != expected {
		t.Errorf("Get(k) expected %v, got %v", expected, v)
	}
	v = lm.Map["k"]
	if v != expected {
		t.Errorf("Map[k] expected %v, got %v", expected, v)
	}
}

func TestUpdateValue(t *testing.T) {
	lm := New()
	expected := "updated"
	lm.Add(1, "v1")
	lm.Add(2, "v2")
	lm.Add(1, expected)

	got := lm.Get(1)
	if got != expected {
		t.Errorf("Update failed, expected=%v, got=%v", expected, got)
	}
}

func TestUpdateNoOrderChange(t *testing.T) {
	lm := New()
	lm.Add(1, "v1")
	lm.Add(2, "v2")
	lm.Add(3, "v3")
	e3 := *lm.last
	e2 := *lm.last.Prev()
	e2next := *lm.last.Prev().next
	e2prev := *lm.last.Prev().prev

	lm.Add(2, "update")
	ne2 := *lm.last.Prev()

	if *lm.last != e3 {
		t.Error("last must not be updated, when val updated")
	}
	if ne2 != e2 {
		t.Error("Second element must not change, when val updated")
	}
	if *ne2.next != e2next {
		t.Error("element.next must not change, when val updated")
	}
	if *ne2.prev != e2prev {
		t.Error("element.prev must not change, when val updated")
	}
}

func TestTwoMaps(t *testing.T) {
	lm1 := New()
	lm1.Add(1, "lm1")
	lm2 := New()
	lm2.Add(1, "lm1")

	if &lm1.Map == &lm2.Map {
		t.Error("Two different linkedmaps must use different maps")
	}

	if lm1.last == lm2.last {
		t.Error("Two different linkedmaps must use different elements")
	}
}

func TestOneMapUpdateAnother(t *testing.T) {
	lm1 := New()
	lm1.Add(1, "lm1")
	lm2 := New()
	lm2.Add(1, "lm2")

	if lm1.Get(1) != "lm1" {
		t.Error("Another linkedmap must not update the first map")
	}
}
