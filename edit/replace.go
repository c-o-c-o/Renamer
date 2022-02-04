package edit

import (
	"errors"
	"regexp"
	"renamer/analyze"
	"strings"

	"gopkg.in/yaml.v2"
)

func (e *Editor) ReplaceName(tinfo analyze.TalkInfo) (*analyze.TalkInfo, error) {
	for _, v := range e.Settings.RepName {
		ptn, new, ok := getKVStrs(v)
		if !ok {
			return &tinfo, errors.New("Replace-Name : it wasn't a string")
		}

		matched, err := regexp.MatchString(ptn, tinfo.Name)
		if err != nil {
			return nil, err
		}

		if matched {
			tinfo.Name = new
			return &tinfo, nil
		}
	}
	return &tinfo, nil
}

func (e *Editor) ReplaceResult(tinfo *analyze.TalkInfo) (string, error) {
	tinfo.Raw["name"] = tinfo.Name
	tinfo.Raw["body"] = tinfo.Body
	rsltname := e.Settings.Rslt.Name

	for k, v := range tinfo.Raw {
		rsltname = strings.ReplaceAll(rsltname, "{@"+k+"}", v)
	}

	for _, v := range e.Settings.RepFileName {
		old, new, ok := getKVStrs(v)
		if !ok {
			return "", errors.New("Replace-FileName : it wasn't a string")
		}

		rsltname = strings.ReplaceAll(rsltname, old, new)
	}

	return rsltname, nil
}

func getKVStrs(item yaml.MapItem) (string, string, bool) {
	ptn, ok := item.Key.(string)
	if !ok {
		return "", "", false
	}

	new, ok := item.Value.(string)
	if !ok {
		return "", "", false
	}

	return ptn, new, true
}
