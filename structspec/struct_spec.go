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

package structspec

import (
	"reflect"
	"strings"
	"sync"
)

var (
	locker = new(sync.RWMutex)
	cache  = make(map[reflect.Type]*StructSpec)
)

type StructSpec struct {
	M  map[string]*FieldSpec
	Cm map[string]*FieldSpec
	Fs []*FieldSpec
}

// GetStructSpec get structure tag info
func GetStructSpec(tagName string, t reflect.Type) *StructSpec {
	locker.RLock()
	ss, found := cache[t]
	locker.RUnlock()
	if found {
		return ss
	}

	locker.Lock()
	defer locker.Unlock()
	ss, found = cache[t]
	if found {
		return ss
	}

	ss = &StructSpec{M: make(map[string]*FieldSpec), Cm: make(map[string]*FieldSpec)}
	compileStructSpec(tagName, t, make(map[string]int), nil, ss)
	cache[t] = ss
	return ss
}

// FieldSpec body 标签解析结果
type FieldSpec struct {
	Tag              string // tag
	Name             string // 字段名
	I                int    // 位置
	Index            []int
	Column           string // 对应数据库字段名
	Type             string // 表字段类型，bool、string、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、float、float64、blob(bytes)、enum、json、date、datetime
	OmitEmpty        bool   // 忽略零值
	OmitInsertEmpty  bool   // INSERT 时忽略零值
	OmitReplaceEmpty bool   // REPLACE 时忽略零值
	OmitUpdateEmpty  bool   // UPDATE 时忽略零值
	OnCreateTime     bool   // INSERT/REPLACE 时初始化为当前时间，具体格式根据 Type 决定，如果是数字类型包括 int、int32、int64 等，则是时间戳，否则就是 time.Time 类型
	OnUpdateTime     bool   // 数据变更时修改为当前时间，具体格式根据 Type 决定，这里我推荐数据库自带的时间戳更新功能。
	OnUniqueID       bool   // 新增数据时候，如果字段为空值，而且类型为 uint64，则自动生成唯一 ID，记得务必在 orm.yaml 配置里面为每台机器设置不同的 machine_id，否则生成的ID可能会有冲突
}

func (ss *StructSpec) FieldSpec(name string) *FieldSpec {
	return ss.M[name]
}

func (ss *StructSpec) ColumnSpec(name string) *FieldSpec {
	return ss.Cm[name]
}

func compileStructSpec(tagName string, t reflect.Type, depth map[string]int, index []int, ss *StructSpec) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch {
		case f.PkgPath != "" && !f.Anonymous:
			// Ignore unexported fields.
		case f.Anonymous:
			typ := f.Type
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}

			if typ.Kind() == reflect.Struct {
				compileStructSpec(tagName, typ, depth, append(index, i), ss)
			}
		default:
			fs := &FieldSpec{Name: f.Name, I: i}
			var tag string

			if tagName != "" {
				tag = f.Tag.Get(tagName)
				if tag != "" {
					fs.Tag = tagName
				}
			} else {
				tag = f.Tag.Get("orm")
				if tag == "" {
					tag = f.Tag.Get("redis")
					if tag == "" { // json to back it up
						tag = f.Tag.Get("json")
						if tag == "" {
							tag = f.Name
							fs.Column = f.Name
						} else {
							fs.Tag = "json"
						}
					} else {
						fs.Tag = "redis"
					}
				} else {
					fs.Tag = "orm"
				}
			}

			p := strings.Split(tag, ",")
			if len(p) > 0 {
				if p[0] == "-" {
					continue
				}

				if len(p[0]) > 0 {
					fs.Column = p[0]
				}

				if len(p) > 1 {
					isOmitOn := passFieldSpec(p[1], fs)
					if !isOmitOn {
						fs.Type = p[1]
					}
				}

				if len(p) > 2 {
					for _, s := range p[2:] {
						_ = passFieldSpec(s, fs)
					}
				}
			}

			d, found := depth[fs.Name]
			if !found {
				d = 1 << 30
			}
			switch {
			case len(index) == d:
				// At same depth, remove from result.
				delete(ss.M, fs.Name)
				j := 0
				for k := 0; k < len(ss.Fs); k++ {
					if fs.Name != ss.Fs[k].Name {
						ss.Fs[j] = ss.Fs[k]
						j++
					}
				}
				ss.Fs = ss.Fs[:j]
			case len(index) < d:
				fs.Index = make([]int, len(index)+1)
				copy(fs.Index, index)
				fs.Index[len(index)] = i
				depth[fs.Name] = len(index)
				ss.M[fs.Name] = fs
				ss.Cm[fs.Column] = fs
				ss.Fs = append(ss.Fs, fs)
			}
		}
	}
}

func passFieldSpec(s string, fs *FieldSpec) bool {
	switch s {
	case "omitempty":
		fs.OmitEmpty = true
		return true
	case "omitinsertempty":
		fs.OmitInsertEmpty = true
		return true
	case "omitreplaceempty":
		fs.OmitReplaceEmpty = true
		return true
	case "omitupdateempty":
		fs.OmitUpdateEmpty = true
		return true
	case "oncreatetime":
		fs.OnCreateTime = true
		return true
	case "onupdatetime":
		fs.OnUpdateTime = true
		return true
	case "onuniqueid":
		fs.OnUniqueID = true
		return true
	default:
		return false
	}
}
