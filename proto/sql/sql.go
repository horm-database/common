package sql

// Join MySQL 表 JOIN
type Join struct {
	Type  string            `json:"type,omitempty"`
	Table string            `json:"table,omitempty"`
	Using []string          `json:"using,omitempty"`
	On    map[string]string `json:"on,omitempty"`
}
