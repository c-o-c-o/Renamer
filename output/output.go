package output

import (
	"os"
	"path/filepath"
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

func FixTextFile(path string, body string) error {
	err := os.WriteFile(path, []byte(body), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
