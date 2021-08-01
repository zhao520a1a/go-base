package xnet

import "net/http"

// 获取 opuser
func GetOpUser(r *http.Request) string {
	var opuser string

	cooku, _ := r.Cookie("user")
	if cooku == nil {
		opuser, _, _ = r.BasicAuth()
	} else {
		opuser = cooku.Value
	}
	return opuser
}
