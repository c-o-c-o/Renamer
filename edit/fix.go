package edit

import "strings"

func (e *Editor) FixBody(body string) string {

	return e.delPreSuf(body)
}

func (e *Editor) delPreSuf(body string) string {
	bdyidx := 0
	bdylen := len(body)

	for _, dp := range e.Settings.DelPrefix {
		if strings.HasPrefix(body, dp) {
			bdyidx = len(dp)
			break
		}
	}

	for _, ds := range e.Settings.DelSuffix {
		if strings.HasSuffix(body, ds) {
			bdylen -= len(ds)
			break
		}
	}

	return body[bdyidx:bdylen]
}
