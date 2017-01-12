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

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var c int
var l *List
var i *Item

func TestMain(t *testing.T) {

	l = NewList()

	Convey("Get with nothing in list", t, func() {
		So(l.Get(3), ShouldBeNil)
	})

	Convey("Can set 2nd item", t, func() {
		l.Put(2, []byte{2})
		i = l.Get(2)
		So(l.Min(), ShouldEqual, l.Get(2))
		So(l.Max(), ShouldEqual, l.Get(2))
		So(i.Ver(), ShouldEqual, 2)
		So(i.Val(), ShouldResemble, []byte{2})
	})

	Convey("Can set 4th item", t, func() {
		l.Put(4, []byte{4})
		i = l.Get(4)
		So(l.Min(), ShouldEqual, l.Get(2))
		So(l.Max(), ShouldEqual, l.Get(4))
		So(i.Ver(), ShouldEqual, 4)
		So(i.Val(), ShouldResemble, []byte{4})
	})

	Convey("Can set 1st item", t, func() {
		l.Put(1, []byte{1})
		i = l.Get(1)
		So(l.Min(), ShouldEqual, l.Get(1))
		So(l.Max(), ShouldEqual, l.Get(4))
		So(i.Ver(), ShouldEqual, 1)
		So(i.Val(), ShouldResemble, []byte{1})
	})

	Convey("Can set 3rd item", t, func() {
		l.Put(3, []byte{3})
		i = l.Get(3)
		So(l.Min(), ShouldEqual, l.Get(1))
		So(l.Max(), ShouldEqual, l.Get(4))
		So(i.Ver(), ShouldEqual, 3)
		So(i.Val(), ShouldResemble, []byte{3})
	})

	Convey("Can set 5th item", t, func() {
		l.Put(5, []byte{5})
		i = l.Get(5)
		So(l.Min(), ShouldEqual, l.Get(1))
		So(l.Max(), ShouldEqual, l.Get(5))
		So(i.Ver(), ShouldEqual, 5)
		So(i.Val(), ShouldResemble, []byte{5})
	})

	Convey("Can get list size", t, func() {
		So(l.Len(), ShouldEqual, 5)
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can get prev item to 1", t, func() {
		So(l.Prev(1), ShouldBeNil)
	})

	Convey("Can get prev item to 3", t, func() {
		i = l.Prev(3)
		So(i.Ver(), ShouldEqual, 2)
		So(i.Val(), ShouldResemble, []byte{2})
	})

	Convey("Can get next item to 3", t, func() {
		i = l.Next(3)
		So(i.Ver(), ShouldEqual, 4)
		So(i.Val(), ShouldResemble, []byte{4})
	})

	Convey("Can get next item to 5", t, func() {
		So(l.Next(5), ShouldBeNil)
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can get upto item at 0", t, func() {
		So(l.Upto(0), ShouldBeNil)
	})

	Convey("Can get upto item at 1", t, func() {
		So(l.Upto(1), ShouldEqual, l.Get(1))
	})

	Convey("Can get upto item at 3", t, func() {
		So(l.Upto(3), ShouldEqual, l.Get(3))
	})

	Convey("Can get upto item at 5", t, func() {
		So(l.Upto(5), ShouldEqual, l.Get(5))
	})

	Convey("Can get upto item at 7", t, func() {
		So(l.Upto(7), ShouldEqual, l.Get(5))
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can get minimum item", t, func() {
		i = l.Min()
		So(i.Ver(), ShouldEqual, 1)
		So(i.Val(), ShouldResemble, []byte{1})
	})

	Convey("Can get maximum item", t, func() {
		i = l.Max()
		So(i.Ver(), ShouldEqual, 5)
		So(i.Val(), ShouldResemble, []byte{5})
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can delete 1st item", t, func() {
		i = l.Del(1)
		So(i.Ver(), ShouldEqual, 1)
		So(i.Val(), ShouldResemble, []byte{1})
	})

	Convey("Can get minimum item", t, func() {
		i = l.Min()
		So(i.Ver(), ShouldEqual, 2)
		So(i.Val(), ShouldResemble, []byte{2})
	})

	Convey("Can get maximum item", t, func() {
		i = l.Max()
		So(i.Ver(), ShouldEqual, 5)
		So(i.Val(), ShouldResemble, []byte{5})
	})

	Convey("Can get list size", t, func() {
		So(l.Len(), ShouldEqual, 4)
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can delete 5th item", t, func() {
		i = l.Del(5)
		So(i.Ver(), ShouldEqual, 5)
		So(i.Val(), ShouldResemble, []byte{5})
	})

	Convey("Can get minimum item", t, func() {
		i = l.Min()
		So(i.Ver(), ShouldEqual, 2)
		So(i.Val(), ShouldResemble, []byte{2})
	})

	Convey("Can get maximum item", t, func() {
		i = l.Max()
		So(i.Ver(), ShouldEqual, 4)
		So(i.Val(), ShouldResemble, []byte{4})
	})

	Convey("Can get list size", t, func() {
		So(l.Len(), ShouldEqual, 3)
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can delete 3rd item", t, func() {
		i = l.Del(3)
		So(i.Ver(), ShouldEqual, 3)
		So(i.Val(), ShouldResemble, []byte{3})
	})

	Convey("Can get minimum item", t, func() {
		i = l.Min()
		So(i.Ver(), ShouldEqual, 2)
		So(i.Val(), ShouldResemble, []byte{2})
	})

	Convey("Can get maximum item", t, func() {
		i = l.Max()
		So(i.Ver(), ShouldEqual, 4)
		So(i.Val(), ShouldResemble, []byte{4})
	})

	Convey("Can get list size", t, func() {
		So(l.Len(), ShouldEqual, 2)
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can replace 2nd item", t, func() {
		l.Put(2, []byte{'R'})
		i = l.Get(2)
		So(i.Ver(), ShouldEqual, 2)
		So(i.Val(), ShouldResemble, []byte{'R'})
	})

	Convey("Can replace 4th item", t, func() {
		l.Put(4, []byte{'R'})
		i = l.Get(4)
		So(i.Ver(), ShouldEqual, 4)
		So(i.Val(), ShouldResemble, []byte{'R'})
	})

	Convey("Can get list size", t, func() {
		So(l.Len(), ShouldEqual, 2)
	})

	// ----------------------------------------
	// ----------------------------------------
	// ----------------------------------------

	Convey("------------------------------", t, nil)

	Convey("Can walk through the list and exit", t, func() {
		var items []*Item
		l.Walk(func(i *Item) bool {
			items = append(items, i)
			return true
		})
		So(len(items), ShouldEqual, 1)
		So(items[0], ShouldEqual, l.Get(2))
	})

	Convey("Can walk through the list without exiting", t, func() {
		var items []*Item
		l.Walk(func(i *Item) bool {
			items = append(items, i)
			return false
		})
		So(len(items), ShouldEqual, 2)
		So(items[0], ShouldEqual, l.Get(2))
		So(items[1], ShouldEqual, l.Get(4))
	})

}
