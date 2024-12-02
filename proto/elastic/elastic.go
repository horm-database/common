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
