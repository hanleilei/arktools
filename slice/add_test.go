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
	"fmt"
	"github.com/hanleilei/arktools/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	// Add 主要依赖于 internal/slice.Add 来保证正确性
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
			name:      "index at end",
			slice:     []int{123, 100},
			addVal:    233,
			index:     2,
			wantSlice: []int{123, 100, 233},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			addVal:    233,
			index:     0,
			wantSlice: []int{233},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:    "index out of range (too large)",
			slice:   []int{123, 100},
			addVal:  233,
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
		{
			name:      "single element insert at 0",
			slice:     []int{123},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123},
		},
		{
			name:      "single element insert at 1",
			slice:     []int{123},
			addVal:    233,
			index:     1,
			wantSlice: []int{123, 233},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 233, 2)
	fmt.Println(res)
	res, _ = Add[int]([]int{}, 233, 0)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 233, -1)
	fmt.Println(err)
	// Output:
	// [1 2 233 3 4]
	// [233]
	// ekit: 下标超出范围，长度 4, 下标 -1
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
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
