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

// List represents a time-series list.
type List struct {
	size int
	min  *Item
	max  *Item
}

// Get gets a specific item from the list. If the exact item does not
// exist in the list, then a nil value is returned.
func (l *List) Get(ver int64) *Item {

	return l.find(ver, true)

}

// Del deletes a specific item from the list, returning the previous item
// if it existed. If it did not exist, a nil value is returned.
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

// Put inserts a new item into the list, ensuring that the list is sorted
// after insertion. If an item with the same version already exists in the
// list, then the value is updated.
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

// Len returns the total number of items in the list.
func (l *List) Len() int {
	return l.size
}

// Min returns the first item in the list. In a time-series list this can be
// used to get the initial version.
func (l *List) Min() *Item {
	return l.min
}

// Max returns the last item in the list. In a time-series list this can be
// used to get the latest version.
func (l *List) Max() *Item {
	return l.max
}

// Seek returns the nearest item in the list, where theversion number is
// less than the given version. In a time-series list, this can be used to
// get the version that was valid at the specified time.
func (l *List) Seek(ver int64) *Item {
	return l.find(ver, false)
}

// Walk iterates over the list starting at the first version, and continuing
// until the walk function returns true.
func (l *List) Walk(fn func(*Item) bool) {
	for i := l.min; !fn(i); i = i.next {
		continue
	}
}

// ---------------------------------------------------------------------------

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
