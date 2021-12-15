package analyze

import (
	"errors"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func GetConverter(enc string) (encoding.Encoding, error) {
	encode := map[string]encoding.Encoding{
		"":          japanese.ShiftJIS,
		"utf-8":     nil,
		"shift-jis": japanese.ShiftJIS,
	}

	v, exist := encode[enc]
	if !exist {
		return nil, errors.New("the encoding was not found")
	}

	return v, nil
}

func Convert(tbyte []byte, conv transform.Transformer) ([]byte, error) {
	ebyte, _, err := transform.Bytes(conv, tbyte)
	if err != nil {
		return nil, err
	}

	return ebyte, nil
}
