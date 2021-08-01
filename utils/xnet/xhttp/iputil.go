package xhttp

import (
	"github.com/zhao520a1a/go-base.git/utils/xnet"
	"net/http"
	"strings"
)

// 获取http请求的client的地址
func IpAddressHttpClient(r *http.Request) string {
	hdr := r.Header
	hdrRealIp := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")

	if hdrRealIp == "" && hdrForwardedFor == "" {
		return xnet.IpAddrFromRemoteAddr(r.RemoteAddr)
	}

	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		for _, ip := range parts {
			ok, _ := xnet.IsInterIp(ip)
			if !ok && len(ip) > 5 && "127." != ip[:4] {
				return ip
			}
		}

	}

	return hdrRealIp
}
