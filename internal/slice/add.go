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

import "github.com/hanleilei/arktools/internal/errs"

// Add 在切片的指定位置插入一个元素，并返回新的切片
// 如果 index 超出范围（< 0 或 > len(src)），返回错误
func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	// 创建新切片，容量为 length+1
	result := make([]T, length+1)

	// 复制 index 之前的元素
	copy(result[:index], src[:index])

	// 插入新元素
	result[index] = element

	// 复制 index 之后的元素
	copy(result[index+1:], src[index:])

	return result, nil
}
