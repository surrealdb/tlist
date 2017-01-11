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

type Item struct {
	ver  int64
	val  []byte
	prev *Item
	next *Item
}

func (i *Item) Del() {
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

func (i *Item) Ver() int64 {
	return i.ver
}

func (i *Item) Val() []byte {
	return i.val
}
