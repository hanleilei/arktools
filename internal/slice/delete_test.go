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
	"github.com/hanleilei/arktools/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			index:     0,
			wantSlice: []int{100},
			wantVal:   123,
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			index:     1,
			wantSlice: []int{123, 125},
			wantVal:   124,
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
			index:     5,
			wantSlice: []int{123, 100, 101, 102, 102},
			wantVal:   102,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			index:   0,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name:      "single element",
			slice:     []int{123},
			index:     0,
			wantSlice: []int{},
			wantVal:   123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestDeleteString(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []string
		index     int
		wantSlice []string
		wantVal   string
		wantErr   error
	}{
		{
			name:      "delete at beginning",
			slice:     []string{"hello", "world"},
			index:     0,
			wantSlice: []string{"world"},
			wantVal:   "hello",
		},
		{
			name:      "delete in middle",
			slice:     []string{"a", "b", "c"},
			index:     1,
			wantSlice: []string{"a", "c"},
			wantVal:   "b",
		},
		{
			name:      "delete at end",
			slice:     []string{"foo", "bar"},
			index:     1,
			wantSlice: []string{"foo"},
			wantVal:   "bar",
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
			res, val, err := Delete(tc.slice, tc.index)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestDeleteStruct(t *testing.T) {
	alice := testutil.Person{Name: "Alice", Age: 30}
	bob := testutil.Person{Name: "Bob", Age: 25}
	david := testutil.Person{Name: "David", Age: 40}
	test := testutil.Person{Name: "Test", Age: 20}

	testCases := []struct {
		name      string
		slice     []testutil.Person
		index     int
		wantSlice []testutil.Person
		wantVal   testutil.Person
		wantErr   error
	}{
		{
			name:      "delete at beginning",
			slice:     []testutil.Person{alice, bob},
			index:     0,
			wantSlice: []testutil.Person{bob},
			wantVal:   alice,
		},
		{
			name: "delete in middle",
			slice: []testutil.Person{
				alice,
				bob,
				david,
			},
			index:     1,
			wantSlice: []testutil.Person{alice, david},
			wantVal:   bob,
		},
		{
			name:      "delete at end",
			slice:     []testutil.Person{alice},
			index:     0,
			wantSlice: []testutil.Person{},
			wantVal:   alice,
		},
		{
			name:    "index out of range",
			slice:   []testutil.Person{test},
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(1, 5),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}
