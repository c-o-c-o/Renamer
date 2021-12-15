package data

import "gopkg.in/yaml.v2"

type Settings struct {
	Rslt struct {
		Name string `yaml:"Name"`
		Enc  string `yaml:"Encode"`
	} `yaml:"Result"`
	Ptns     []Pattern     `yaml:"Patterns"`
	DelPrefs []string      `yaml:"Delete-Prefix"`
	RepName  yaml.MapSlice `yaml:"Replace-Name"`
}

type Pattern struct {
	Tgt string `yaml:"Target"`
	Ptn string `yaml:"Pattern"`
	Enc string `yaml:"Encode"`
}
