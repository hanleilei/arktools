// Copyright 2026 hanleilei
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"testing"

	"github.com/hanleilei/arktools/internal/errs"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			addVal:    233,
			index:     1,
			wantSlice: []int{123, 233, 124, 125},
		},
		{
			name:    "index out of range",
			slice:   []int{123, 100},
			index:   12,
			wantErr: errs.NewErrIndexOutOfRange(2, 12),
		},
		{
			name:    "index less than 0",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			addVal:    233,
			index:     5,
			wantSlice: []int{123, 100, 101, 102, 102, 233, 102},
		},
		{
			name:      "append on last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			addVal:    233,
			index:     6,
			wantSlice: []int{123, 100, 101, 102, 102, 102, 233},
		},
		{
			name:    "index out of range",
			slice:   []int{123, 100, 101, 102, 102, 102},
			addVal:  233,
			index:   7,
			wantErr: errs.NewErrIndexOutOfRange(6, 7),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func TestAddString(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []string
		addVal    string
		index     int
		wantSlice []string
		wantErr   error
	}{
		{
			name:      "insert at beginning",
			slice:     []string{"hello", "world"},
			addVal:    "hi",
			index:     0,
			wantSlice: []string{"hi", "hello", "world"},
		},
		{
			name:      "insert in middle",
			slice:     []string{"a", "b", "c"},
			addVal:    "x",
			index:     1,
			wantSlice: []string{"a", "x", "b", "c"},
		},
		{
			name:      "insert at end",
			slice:     []string{"foo", "bar"},
			addVal:    "baz",
			index:     2,
			wantSlice: []string{"foo", "bar", "baz"},
		},
		{
			name:    "index out of range",
			slice:   []string{"test"},
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(1, 5),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

type Person struct {
	Name string
	Age  int
}

func TestAddStruct(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []Person
		addVal    Person
		index     int
		wantSlice []Person
		wantErr   error
	}{
		{
			name: "insert at beginning",
			slice: []Person{
				{Name: "Alice", Age: 30},
				{Name: "Bob", Age: 25},
			},
			addVal: Person{Name: "Charlie", Age: 35},
			index:  0,
			wantSlice: []Person{
				{Name: "Charlie", Age: 35},
				{Name: "Alice", Age: 30},
				{Name: "Bob", Age: 25},
			},
		},
		{
			name: "insert in middle",
			slice: []Person{
				{Name: "Alice", Age: 30},
				{Name: "Bob", Age: 25},
				{Name: "David", Age: 40},
			},
			addVal: Person{Name: "Eve", Age: 28},
			index:  1,
			wantSlice: []Person{
				{Name: "Alice", Age: 30},
				{Name: "Eve", Age: 28},
				{Name: "Bob", Age: 25},
				{Name: "David", Age: 40},
			},
		},
		{
			name: "insert at end",
			slice: []Person{
				{Name: "Alice", Age: 30},
			},
			addVal: Person{Name: "Frank", Age: 50},
			index:  1,
			wantSlice: []Person{
				{Name: "Alice", Age: 30},
				{Name: "Frank", Age: 50},
			},
		},
		{
			name:    "index out of range",
			slice:   []Person{{Name: "Test", Age: 20}},
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(1, 5),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
