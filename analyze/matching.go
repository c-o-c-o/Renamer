package analyze

import (
	"errors"
	"regexp"
	"renamer/data"
	"strings"
)

type TalkInfo struct {
	Name string
	Body string
	Raw  map[string]string
}

func GetTalkInfo(tgts *Targets, ptns []data.Pattern) (*TalkInfo, error) {
	for _, ptn := range ptns {
		ttext, err := GetTargetValue(ptn.Tgt, tgts.FileDir, tgts.FileName, string(tgts.FileBody))
		if err != nil {
			return nil, err
		}

		ttext = removeNewLine(ttext)

		hitmap, err := SubmatchMap(ttext, ptn.Ptn)
		if err != nil {
			return nil, err
		}
		if hitmap == nil {
			continue
		}

		return &TalkInfo{
			Name: hitmap["name"],
			Body: hitmap["body"],
			Raw:  hitmap,
		}, nil
	}

	return nil, errors.New("did not match all patterns")
}

func removeNewLine(ttext string) string {
	ttext = strings.ReplaceAll(ttext, "\n", "")
	ttext = strings.ReplaceAll(ttext, "\r", "")
	return ttext
}

func SubmatchMap(str string, ptn string) (map[string]string, error) {
	reg, err := regexp.Compile(ptn)
	if err != nil {
		return nil, err
	}

	match := reg.FindStringSubmatch(str)
	if match == nil {
		return nil, nil
	}

	rslt := make(map[string]string)
	for i, name := range reg.SubexpNames() {
		if i != 0 && name != "" {
			rslt[name] = match[i]
		}
	}

	return rslt, nil
}

func GetTargetValue(tgt string, dir string, name string, text string) (string, error) {

	return strings.NewReplacer(
		"Dir", dir,
		"Name", name,
		"Text", text).Replace(tgt), nil
}
