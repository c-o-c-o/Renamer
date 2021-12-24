package edit

import (
	"errors"
	"regexp"
	"renamer/analyze"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReplaceTalkInfo(tinfo analyze.TalkInfo, repname yaml.MapSlice) (*analyze.TalkInfo, error) {
	for _, v := range repname {
		ptn, ok := v.Key.(string)
		if !ok {
			return nil, errors.New("Replace-Name : it wasn't a string")
		}

		new, ok := v.Value.(string)
		if !ok {
			return nil, errors.New("Replace-Name : it wasn't a string")
		}

		matched, err := regexp.MatchString(ptn, tinfo.Name)
		if err != nil {
			return nil, err
		}

		if matched {
			tinfo.Name = new
			tinfo.Raw["name"] = new
			break
		}
	}

	return &tinfo, nil
}

func DeletePrefix(body string, delprefs []string) string {
	for _, dp := range delprefs {
		if len(dp) < len(body) && strings.HasPrefix(body, dp) {
			return body[len(dp):]
		}
	}
	return body
}

func ReplaceResult(rsltname string, tinfo *analyze.TalkInfo) string {
	invrep := map[string]string{
		"\\": "￥",
		"/":  "／",
		":":  "：",
		"*":  "＊",
		"?":  "？",
		"\"": "”",
		"<":  "＜",
		">":  "＞",
		"|":  "｜",
	}

	for k, v := range tinfo.Raw {
		rsltname = strings.ReplaceAll(rsltname, "{@"+k+"}", v)
	}

	for k, v := range invrep {
		rsltname = strings.ReplaceAll(rsltname, k, v)
	}

	return rsltname
}
