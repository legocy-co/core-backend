package helpers

import "net/url"

func EncodeURLString(s string) string {
	return url.QueryEscape(s)
}
