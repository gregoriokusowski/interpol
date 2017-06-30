package interpol

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type locale struct {
	Name  string   `yaml:"name"`
	Files []string `yaml:"files"`
}

type config struct {
	Master  locale   `yaml:"master"`
	Locales []locale `yaml:"locales"`
}

func parseConfig(fileName string) config {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var c config
	yaml.Unmarshal(file, &c)
	return c
}
