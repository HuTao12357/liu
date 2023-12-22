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

// StringAdd 字符串拼接
func StringAdd(arr string, arrAdd string) string {
	byte1 := []byte(arr)
	byte2 := []byte(arrAdd)
	byte1 = append(byte1, byte2...)
	arr = string(byte1)
	return arr
}
