package analyze

import (
	"os"
	"path/filepath"
)

type Targets struct {
	FileDir  string
	FileName string
	FileBody []byte
}

func GetPtnTgts(textpath string) (*Targets, error) {
	b, err := os.ReadFile(textpath)
	if err != nil {
		return nil, err
	}

	return &Targets{
		FileDir:  filepath.ToSlash(filepath.Dir(textpath)),
		FileName: withoutExt(textpath),
		FileBody: b,
	}, nil
}

func withoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
