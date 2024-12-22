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

package elastic

type ItemInfo struct {
	Score  *float64   `json:"_score,omitempty"`  // computed score
	Index  string     `json:"_index,omitempty"`  // index name
	Type   string     `json:"_type,omitempty"`   // type meta field
	Id     string     `json:"_id,omitempty"`     // external or internal
	Nested *NestedHit `json:"_nested,omitempty"` // for nested inner hits
}

// NestedHit is a nested innerhit
type NestedHit struct {
	Field  string     `json:"field"`
	Offset int        `json:"offset,omitempty"`
	Child  *NestedHit `json:"_nested,omitempty"`
}
