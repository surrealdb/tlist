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

const (
	upto int8 = iota
	prev
	next
	exact
	nearest
)

func NewList() *List {
	return &List{}
}

// Put inserts a new item into the list, ensuring that the list is sorted
// after insertion. If an item with the same version already exists in the
// list, then the value is updated.
func (l *List) Put(ver int64, val []byte) {

	// If there is no min or max for
	// this list, then we can just add
	// this item as the min and max.

	if l.min == nil && l.max == nil {
		i := &Item{ver: ver, val: val}
		l.min, l.max = i, i
		l.size++
		return
	}

	// Otherwise find the nearest item
	// to this version so we can update
	// it or prepend / append to it.

	f := l.find(ver, nearest)

	if f.ver == ver {
		f.val = val
		return
	}

	// If the found item version is not
	// the same version as the one we
	// updating then create a new item.

	i := &Item{ver: ver, val: val}

	if f.ver < ver {
		if f.next != nil {
			f.next.prev = i
			i.next = f.next
			f.next = i
		}
		i.prev = f
		f.next = i
	}

	if f.ver > ver {
		i.next = f
		f.prev = i
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

	l.size++

	return

}

// Del deletes a specific item from the list, returning the previous item
// if it existed. If it did not exist, a nil value is returned.
func (l *List) Del(ver int64) *Item {

	i := l.find(ver, exact)

	if i != nil {

		if i.prev != nil && i.next != nil {
			i.prev.next = i.next
			i.next.prev = i.prev
			i.prev = nil
			i.next = nil
		} else if i.prev != nil {
			i.prev.next = nil
			l.max = i.prev
			i.prev = nil
		} else if i.next != nil {
			i.next.prev = nil
			l.min = i.next
			i.next = nil
		}

		l.size--

	}

	return i

}

// Get gets a specific item from the list. If the exact item does not
// exist in the list, then a nil value is returned.
func (l *List) Get(ver int64) *Item {
	return l.find(ver, exact)
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

// Prev returns the nearest item in the list, where the version number is
// less than the given version. In a time-series list, this can be used
// to get the version that was valid before a specified time.
func (l *List) Prev(ver int64) *Item {
	return l.find(ver, prev)
}

// Next returns the nearest item in the list, where the version number is
// greater than the given version. In a time-series list, this can be used
// to get the version that was changed after a specified time.
func (l *List) Next(ver int64) *Item {
	return l.find(ver, next)
}

// Upto returns the nearest item in the list, where the version number is
// less than or equal to the given version. In a time-series list, this can
// be used to get the version that was current at the specified time.
func (l *List) Upto(ver int64) *Item {
	return l.find(ver, upto)
}

// Walk iterates over the list starting at the first version, and continuing
// until the walk function returns true.
func (l *List) Walk(fn func(*Item) bool) {
	for i := l.min; i != nil && !fn(i); i = i.next {
		continue
	}
}

// ---------------------------------------------------------------------------

func (l *List) find(ver int64, what int8) (i *Item) {

	if l.min == nil && l.max == nil {
		return nil
	}

	switch what {

	case prev: // Get the item below the specified version

		if i = l.find(ver, upto); i != nil {
			return i.prev
		}

	case next: // Get the item above the specified version

		if i = l.find(ver, upto); i != nil {
			return i.next
		}

	case upto: // Get the item up to the specified version

		if l.min.ver <= ver {
			for i = l.min; i != nil && i.next != nil && i.ver < ver; i = i.next {
				// Ignore
			}
			return i
		}

	case exact: // Get the exact specified version

		for i = l.min; i != nil && i.ver <= ver; i = i.next {
			if i.ver == ver {
				return i
			}
		}

	case nearest: // Get the item nearest the specified version

		for i = l.min; i != nil; i = i.next {
			if i.ver == ver || i.next == nil || i.next.ver > ver {
				return i
			}
		}

	}

	return nil

}
