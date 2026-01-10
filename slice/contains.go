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

// toMap 将切片转为 map，便于快速查找
// NOTE: contains.go 不再定义 toMap，统一用 map.go 的实现

// Contains 判断 src 里面是否存在 dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == dst
	})
}

// ContainsFunc 判断 src 里面是否存在满足条件的元素
// match: 自定义判断函数，返回 true 表示找到目标
// 推荐优先使用 Contains，复杂场景用 ContainsFunc
func ContainsFunc[T any](src []T, match func(src T) bool) bool {
	// 遍历调用equal函数进行判断
	for _, v := range src {
		if match(v) {
			return true
		}
	}
	return false
}

// ContainsAny 判断 src 里面是否存在 dst 中的任何一个元素
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
// equal: 自定义判断函数，返回 true 表示两个元素相等
// 推荐优先使用 ContainsAny，复杂场景用 ContainsAnyFunc
// 性能优化建议：如需处理大切片，可考虑自定义 key 并用 map 加速查找
func ContainsAnyFunc[T any](src, dst []T, equal func(src, dst T) bool) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断 src 里面是否存在 dst 中的所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断 src 里面是否存在 dst 中的所有元素
// equal: 自定义判断函数，返回 true 表示两个元素相等
// 推荐优先使用 ContainsAll，复杂场景用 ContainsAllFunc
// 性能优化建议：如需处理大切片，可考虑自定义 key 并用 map 加速查找
func ContainsAllFunc[T any](src, dst []T, equal func(src, dst T) bool) bool {
	for _, valDst := range dst {
		if !ContainsFunc[T](src, func(src T) bool {
			return equal(src, valDst)
		}) {
			return false
		}
	}
	return true
}
