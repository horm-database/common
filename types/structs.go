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
	"reflect"
	"strings"
	"sync"

	"github.com/horm-database/common/snowflake"
)

const (
	OpInsert  = 1
	OpReplace = 2
	OpUpdate  = 3
)

// StructToMap 结构体转 map
func StructToMap(v reflect.Value, tag string, op int8) map[string]interface{} {
	ss := GetStructSpec(tag, v.Type())

	data := make(map[string]interface{})

	for _, fs := range ss.M {
		if fs.Tag != tag {
			continue
		}

		iv := v.FieldByIndex(fs.Index)
		isEmpty := IsEmpty(iv)

		if fs.OnUniqueID && op == OpInsert && isEmpty && iv.Kind() == reflect.Uint64 {
			data[fs.Column] = snowflake.GenerateID()
			if fs.EsID {
				data["_id"] = data[fs.Column]
			}
			continue
		}

		if fs.OnUpdateTime && isEmpty {
			//修改时自动赋值当前时间，仅在值为零值时才自动赋值
			data[fs.Column] = GetFormatTime(Now(fs.Type), fs.TimeFmt)
			continue
		} else if (op == OpInsert || op == OpReplace) && fs.OnCreateTime && isEmpty {
			//自动插入当前时间，仅在值为零值时才自动赋值
			data[fs.Column] = GetFormatTime(Now(fs.Type), fs.TimeFmt)
			continue
		}

		if isEmpty && (fs.OmitEmpty ||
			(op == OpInsert && fs.OmitInsertEmpty) ||
			(op == OpReplace && fs.OmitReplaceEmpty) ||
			(op == OpUpdate && fs.OmitUpdateEmpty)) { // 忽略零值
			continue
		}

		data[fs.Column] = getValue(fs, iv)

		if op == OpInsert && fs.EsID {
			data["_id"] = data[fs.Column]
		}
	}

	return data
}

// StructsToMaps 结构体数组转 map 数组
func StructsToMaps(arrV reflect.Value, tag string, isInsert bool) []map[string]interface{} {
	arrLen := arrV.Len() //数组长度
	if arrLen <= 0 {
		return nil
	}

	ss := GetStructSpec(tag, reflect.Indirect(arrV.Index(0)).Type())

	ignores := getIgnores(ss, arrV, arrLen, isInsert)

	datas := []map[string]interface{}{}

	//插入语句
	for k := 0; k < arrLen; k++ {
		kv := reflect.Indirect(arrV.Index(k))

		data := map[string]interface{}{}

		for name, fs := range ss.M {
			if fs.Tag != "orm" {
				continue
			}

			iv := kv.FieldByIndex(fs.Index)
			isEmpty := IsEmpty(iv)

			if ignore := ignores[name]; !ignore {
				//自动插入当前时间，仅在值为零值时才自动赋值
				if (fs.OnCreateTime || fs.OnUpdateTime) && isEmpty {
					data[fs.Column] = GetFormatTime(Now(fs.Type), fs.TimeFmt)
				} else if fs.OnUniqueID && isInsert && isEmpty && iv.Kind() == reflect.Uint64 {
					data[fs.Column] = snowflake.GenerateID()
				} else {
					data[fs.Column] = getValue(fs, iv)
				}

				if isInsert && fs.EsID {
					data["_id"] = data[fs.Column]
				}
			} else if fs.OnUniqueID && isInsert && isEmpty && iv.Kind() == reflect.Uint64 {
				data[fs.Column] = snowflake.GenerateID()
			} else if (fs.OnCreateTime || fs.OnUpdateTime) && isEmpty {
				data[fs.Column] = GetFormatTime(Now(fs.Type), fs.TimeFmt)
			}
		}

		datas = append(datas, data)
	}

	return datas
}

func getValue(fs *FieldSpec, iv reflect.Value) interface{} {
	if !iv.CanInterface() {
		return nil
	}

	return GetFormatTime(iv.Interface(), fs.TimeFmt)
}

// getIgnores 获取忽略字段
func getIgnores(ss *StructSpec,
	arrV reflect.Value, arrLen int, isInsert bool) map[string]bool {
	//获取忽略字段
	ignores := map[string]bool{}
	for name, fs := range ss.M {
		if fs.OmitEmpty || (isInsert && fs.OmitInsertEmpty) || (!isInsert && fs.OmitReplaceEmpty) {
			ignores[name] = true
		} else {
			ignores[name] = false
		}
	}

	for k := 0; k < arrLen; k++ {
		kv := reflect.Indirect(arrV.Index(k))

		for name := range ss.M {
			if ignore := ignores[name]; ignore {
				iv := kv.FieldByName(name)
				if !IsEmpty(iv) { // 存在非空值，则该字段不忽略
					ignores[name] = false
				}
			}
		}
	}

	return ignores
}

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
	Type             Type   // orm 类型，不同数据库会映射到不同类型
	OmitEmpty        bool   // 忽略零值
	OmitInsertEmpty  bool   // INSERT 时忽略零值
	OmitReplaceEmpty bool   // REPLACE 时忽略零值
	OmitUpdateEmpty  bool   // UPDATE 时忽略零值
	OnCreateTime     bool   // INSERT/REPLACE 时初始化为当前时间，具体格式根据 Type 决定，如果是数字类型包括 int、int32、int64 等，则是时间戳，否则就是 time.Time 类型，如果设置了该属性，则在插入数据时 omit empty 属性失效。
	OnUpdateTime     bool   // 数据变更时修改为当前时间，具体格式根据 Type 决定，这里我推荐数据库自带的时间戳更新功能，如果设置了该属性，则在插入/修改数据时 omit empty 属性失效。
	TimeFmt          string // 当字段底层类型为 time.Time 时，格式化时间，仅针对请求格式化，返回数据的解析在 codec 内。
	OnUniqueID       bool   // 新增数据时候，如果字段为空值，而且类型为 uint64，则自动生成唯一 ID，如果设置了该属性，则在插入数据时 omit empty 属性失效，记得务必在 orm.yaml 配置里面为每台机器设置不同的 machine_id，否则生成的ID可能会有冲突。
	EsID             bool   // 是否 es 主键 _id
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
						fs.Type = OrmType[strings.ToLower(p[1])]
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
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "time_fmt") {
		fs.TimeFmt = strings.Trim(strings.TrimPrefix(s, "time_fmt="), "'")
	}

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
	case "es_id":
		fs.EsID = true
		return true
	default:
		return false
	}
}
