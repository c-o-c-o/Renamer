package edit

import "strings"

func FixBody(body string, delprefix []string, delsuffix []string) string {
	bdyidx := 0
	bdylen := len(body)

	for _, dp := range delprefix {
		if strings.HasPrefix(body, dp) {
			bdyidx = len(dp)
			break
		}
	}

	for _, ds := range delsuffix {
		if strings.HasSuffix(body, ds) {
			bdylen -= len(ds)
			break
		}
	}

	return body[bdyidx:bdylen]
}
