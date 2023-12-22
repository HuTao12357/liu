package util

import "strings"

type ErrorMsg struct {
	msg string
}

func (e ErrorMsg) Errors() string {
	return e.msg
}

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

// StringSub 字符串截取(index)
func StringSub(arr string, begin int, end int) (string, error) {
	if end > len(arr) || begin < 0 || end < 0 {
		err := &ErrorMsg{"索引越界"}
		panic(err.msg)
	}
	arr2 := arr[begin:end]
	return arr2, nil
}

// StringReplaceFirst 字符串替换 (first)
func StringReplaceFirst(arr string, old string, new string) string {
	str := strings.Replace(arr, old, new, 1)
	return str
}

// StringReplaceAll 字符串替换 (All)
func StringReplaceAll(arr string, old string, new string) string {
	str := strings.Replace(arr, old, new, -1)
	return str
}
