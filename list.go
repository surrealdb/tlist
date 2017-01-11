// Copyright © 2016 Abcum Ltd
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

package tlist

type List struct {
	size int
	min  *Item
	max  *Item
}

func (l *List) Get(ver int64) *Item {

	return l.find(ver, true)

}

func (l *List) Del(ver int64) *Item {

	i := l.find(ver, true)

	if i != nil {
		if i.prev != nil && i.next != nil {
			i.prev.next = i.next
			i.next.prev = i.prev
			i.prev = nil
			i.next = nil
		} else if i.prev != nil {
			i.prev.next = nil
			i.prev = nil
		} else if i.next != nil {
			i.next.prev = nil
			i.next = nil
		}
	}

	return i

}

func (l *List) Put(ver int64, val []byte) {

	// If the exact version item already
	// exists, then set its value and
	// return it immediately.

	e := l.find(ver, true)

	if e != nil {
		e.val = val
		return
	}

	// Otherwise find the nearest item
	// and insert this item after it
	// and before any next item.

	i := &Item{ver: ver, val: val}

	f := l.find(ver, false)

	if f != nil {
		if f.next != nil {
			f.next.prev = i
			i.next = f.next
			i.prev = f
			f.next = i
		} else {
			i.prev = f
			f.next = i
		}
	}

	// If there are no previous items
	// before this item then mark this
	// item as the minimum in the list.

	if i.prev == nil {
		l.min = i
	}

	// If there are no subsequent items
	// after this item then mark this
	// item as the maximum in the list.

	if i.next == nil {
		l.max = i
	}

	return

}

func (l *List) Len() int {
	return l.size
}

func (l *List) Min() *Item {
	return l.min
}

func (l *List) Max() *Item {
	return l.max
}

func (l *List) Seek(ver int64) *Item {
	return l.find(ver, false)
}

func (l *List) Walk(fn func(*Item) bool) {
	for i := l.min; !fn(i); i = i.next {
		continue
	}
}

func (l *List) find(ver int64, exact bool) *Item {

	if exact == true {
		for i := l.min; i != nil; i = i.next {
			if i.ver == ver {
				return i
			}
		}
	}

	if exact == false {
		for i := l.min; i != nil; i = i.next {
			if i.next != nil && i.next.ver > ver {
				return i
			}
		}
	}

	return nil

}
