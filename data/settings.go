package data

import "gopkg.in/yaml.v2"

type Settings struct {
	Rslt struct {
		Name string `yaml:"Name"`
		Enc  string `yaml:"Encode"`
	} `yaml:"Result"`
	Ptns        []Pattern     `yaml:"Patterns"`
	DelPrefix   []string      `yaml:"Delete-Prefix"`
	DelSuffix   []string      `yaml:"Delete-Suffix"`
	RepName     yaml.MapSlice `yaml:"Replace-Name"`
	RepFileName yaml.MapSlice `yaml:"Replace-FileName"`
}

type Pattern struct {
	Tgt string `yaml:"Target"`
	Ptn string `yaml:"Pattern"`
	Enc string `yaml:"Encode"`
}
