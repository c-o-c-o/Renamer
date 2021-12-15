package output

import (
	"os"
	"path/filepath"
	"renamer/analyze"
)

func Renames(nname string, rnpaths []string) error {
	var err error

	for _, p := range rnpaths {
		err = os.Rename(p, filepath.Join(filepath.Dir(p), nname+filepath.Ext(p)))
		if err != nil {
			println(err.Error())
		}
	}

	return err
}

func FixTextFile(path string, body string, enc string) error {
	conv, err := analyze.GetConverter(enc)
	if err != nil {
		return err
	}

	bbody := []byte(body)
	if conv != nil {
		bbody, err = analyze.Convert(bbody, conv.NewEncoder())
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(path, bbody, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
