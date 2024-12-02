package types

// HasBytes 数组中是否有 bytes
func HasBytes(arr []interface{}) bool {
	for _, v := range arr {
		switch v.(type) {
		case []byte, *[]byte:
			return true
		}
	}

	return false
}
