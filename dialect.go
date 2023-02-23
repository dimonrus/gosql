package gosql

import (
	"strconv"
	"unsafe"
)

// QueryHook allows to apply custom hook after build query
var QueryHook func(query string) string

// PostgresQueryParamHook Position argument
func PostgresQueryParamHook(query string) string {
	var ll = len(query)
	var b = make([]byte, ll*2)
	var j int64 = 1
	var i, k, l, s int
	for i < len(query) {
		if query[i] == '?' {
			p := query[s:i] + "$" + strconv.FormatInt(j, 10)
			l += len(p)
			if l > len(b) {
				ll = ll * 2
				b = append(b, make([]byte, ll)...)
			}
			copy(b[k:l], p)
			s = i + 1
			k = l
			j++
		}
		i++
	}
	if i > s {
		p := query[s:]
		l += len(p)
		copy(b[k:l], p)
	}
	b = b[:l]
	return *(*string)(unsafe.Pointer(&b))
}
