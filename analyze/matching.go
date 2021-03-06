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
		conv, err := GetConverter(ptn.Enc)
		if err != nil {
			return nil, err
		}

		fbodyb := tgts.FileBody
		if conv != nil {
			fbodyb, err = Convert(tgts.FileBody, conv.NewDecoder())
			if err != nil {
				return nil, err
			}
		}

		ttext, err := GetTargetValue(ptn.Tgt, tgts.FileName, string(fbodyb))
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

func GetTargetValue(tgt string, name string, text string) (string, error) {
	tgts := strings.Split(tgt, ":")
	r := make([]string, 0, len(tgts))
	t := map[string]string{
		"Name": name,
		"Text": text,
	}

	for _, v := range tgts {
		s, exist := t[v]
		if !exist {
			return "", errors.New("the target was not found")
		}

		r = append(r, s)
	}

	return strings.Join(r, ":"), nil
}
