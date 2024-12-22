// Copyright (c) 2024 The horm-database Authors. All rights reserved.
// This file Author:  CaoHao <18500482693@163.com> .
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func SplitInt(s, sep string) []int {
	s = strings.TrimSpace(s)
	if s == "" {
		return []int{}
	}

	tmp := strings.Split(s, sep)
	ret := make([]int, len(tmp))
	for k, v := range tmp {
		ret[k], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	return ret
}

func SplitInt8(s, sep string) []int8 {
	s = strings.TrimSpace(s)
	if s == "" {
		return []int8{}
	}

	tmp := strings.Split(s, sep)
	ret := make([]int8, len(tmp))
	for k, v := range tmp {
		i, _ := strconv.Atoi(strings.TrimSpace(v))
		ret[k] = int8(i)
	}

	return ret
}

func SplitInt64(s, sep string) []int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return []int64{}
	}

	tmp := strings.Split(s, sep)
	ret := make([]int64, len(tmp))
	for k, v := range tmp {
		ret[k], _ = strconv.ParseInt(strings.TrimSpace(v), 10, 64)
	}

	return ret
}

func SplitUint64(s, sep string) []uint64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return []uint64{}
	}

	tmp := strings.Split(s, sep)
	ret := make([]uint64, len(tmp))
	for k, v := range tmp {
		ret[k], _ = strconv.ParseUint(strings.TrimSpace(v), 10, 64)
	}

	return ret
}

func JoinUint64(s []uint64, sep string) string {
	if len(s) == 0 {
		return ""
	}

	var ret strings.Builder

	for k, v := range s {
		if k > 0 {
			ret.WriteString(sep)
		}

		ret.WriteString(strings.TrimSpace(fmt.Sprint(v)))
	}

	return ret.String()
}

func JoinInt8(s []int8, sep string) string {
	if len(s) == 0 {
		return ""
	}

	var ret strings.Builder

	for k, v := range s {
		if k > 0 {
			ret.WriteString(sep)
		}

		ret.WriteString(strings.TrimSpace(fmt.Sprint(v)))
	}

	return ret.String()
}

// CutString 以 sep 作为分隔符切割字符串
func CutString(s, sep string) (found bool, s1, s2 string) {
	if s == sep {
		return true, "", ""
	}

	i := strings.Index(s, sep)

	switch i {
	case -1:
		found = false
		s2 = s
	default:
		found = true
		s1 = s[:i]
		s2 = s[i+len(sep):]
	}

	return
}

// FirstWord 开始 n 个字符
func FirstWord(key string, n int) string {
	if key == "" || n <= 0 {
		return ""
	}

	if n >= len(key) {
		return key
	}

	return key[0:n]
}

// CutLast 移除最后 n 个字符
func CutLast(key string, n int) string {
	if len(key) <= n {
		return ""
	}

	return key[0 : len(key)-n]
}

// LastWord 末尾 n 个字符
func LastWord(key string, n int) string {
	if key == "" || n <= 0 {
		return ""
	}

	l := len(key)

	if n >= l {
		return key
	}

	return key[l-n:]
}

// QuickReplaceLFCR 替换 \r(回车)、\n(换行) 为字符串 `\r`、`\n`
func QuickReplaceLFCR(b []byte) string {
	var j, k int
	var indexs []int

	var l = len(b)

	for {
		c, size := utf8.DecodeRune(b[j:])

		if c == '\r' || c == '\n' {
			indexs = append(indexs, j)
		}

		j += size
		if j >= l {
			break
		}
	}

	if len(indexs) > 0 {
		var ret = make([]byte, len(b)+len(indexs))

		j = 0

		for _, index := range indexs {
			if index != j {
				copy(ret[k:], b[j:index])
				k += index - j
				j = index
			}

			if b[j] == '\r' {
				copy(ret[k:], "\\r")
			} else {
				copy(ret[k:], "\\n")
			}
			k += 2
			j++
		}

		if j < l {
			copy(ret[k:], b[j:])
		}
		return BytesToString(ret)
	} else {
		return BytesToString(b)
	}
}

// QuickReplaceLFCR2Space 替换 \r(回车)、\n(换行) 为空格
func QuickReplaceLFCR2Space(b []byte) string {
	var j int

	var l = len(b)

	for {
		c, size := utf8.DecodeRune(b[j:])

		if c == '\r' || c == '\n' {
			b[j] = ' '
		}

		j += size
		if j >= l {
			break
		}
	}

	return BytesToString(b)
}

// QuickRemoveLFCR 去掉 \r(回车)、\n(换行)
func QuickRemoveLFCR(b []byte) string {
	var j, k int
	var indexs []int

	var l = len(b)

	for {
		c, size := utf8.DecodeRune(b[j:])

		if c == '\r' || c == '\n' {
			indexs = append(indexs, j)
		}

		j += size
		if j >= l {
			break
		}
	}

	if len(indexs) > 0 {
		var ret = make([]byte, len(b)-len(indexs))

		j = 0

		for _, index := range indexs {
			if index != j {
				copy(ret[k:], b[j:index])
				k += index - j
				j = index
			}

			j++
		}

		if j < l {
			copy(ret[k:], b[j:])
		}
		return BytesToString(ret)
	} else {
		return BytesToString(b)
	}
}
