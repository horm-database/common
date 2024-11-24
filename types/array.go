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

// InArrayInt needle 是否在 arr 里面
func InArrayInt(arr []int, needle int) bool {
	for _, v := range arr {
		if needle == v {
			return true
		}
	}
	return false
}

// InArrayInt8 needle 是否在 arr 里面
func InArrayInt8(arr []int8, needle int8) bool {
	for _, v := range arr {
		if needle == v {
			return true
		}
	}
	return false
}

// InArrayUint64 needle 是否在 arr 里面
func InArrayUint64(arr []uint64, needle uint64) bool {
	for _, v := range arr {
		if needle == v {
			return true
		}
	}
	return false
}

// InArrayString needle 是否在 arr 里面
func InArrayString(arr []string, needle string) bool {
	for _, v := range arr {
		if needle == v {
			return true
		}
	}
	return false
}

// UniqUint64 去重
func UniqUint64(arr []uint64) []uint64 {
	if len(arr) == 0 {
		return arr
	}

	tmp := map[uint64]bool{}

	ret := []uint64{}
	for _, v := range arr {
		_, ok := tmp[v]
		if ok {
			continue
		}

		tmp[v] = true
		ret = append(ret, v)
	}

	return ret

}
