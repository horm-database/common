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

package crypto

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/horm-database/common/types"
)

// MD5Str 计算字符串的 md5 值
func MD5Str(str string) string {
	sumBytes := md5.Sum(types.StringToBytes(str))
	return hex.EncodeToString(sumBytes[:])
}

// MD5Bytes 计算字符串的 md5 值
func MD5Bytes(b []byte) []byte {
	src := md5.Sum(b)
	dst := make([]byte, hex.EncodedLen(len(src[:])))
	hex.Encode(dst, src[:])
	return dst
}
