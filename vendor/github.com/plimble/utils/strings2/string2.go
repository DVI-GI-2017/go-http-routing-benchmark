package string2

import (
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Concat(s ...string) string {
	size := 0
	for i := 0; i < len(s); i++ {
		size += len(s[i])
	}

	buf := make([]byte, 0, size)

	for i := 0; i < len(s); i++ {
		buf = append(buf, StringToBytes(s[i])...)
	}

	return BytesToString(buf)
}

func ConcatBase(base string, s ...string) string {
	size := len(base)
	for i := 0; i < len(s); i++ {
		size += len(s[i])
	}

	buf := make([]byte, 0, size)
	buf = append(buf, StringToBytes(base)...)

	for i := 0; i < len(s); i++ {
		buf = append(buf, StringToBytes(s[i])...)
	}

	return BytesToString(buf)
}

// func Concat(s ...string) string {
// 	buf := &bytes.Buffer{}
// 	for i := 0; i < len(s); i++ {
// 		buf.WriteString(s[i])
// 	}

// 	return buf.String()
// }
