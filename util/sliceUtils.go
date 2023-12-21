package util

// IsEmptySliStr 判断切片string是否为空
func IsEmptySliStr(arr []string) bool {
	if arr == nil {
		return false
	}
	if len(arr) == 0 {
		return false
	}
	return true
}
