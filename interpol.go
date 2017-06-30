package interpol

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"

	yaml "gopkg.in/yaml.v2"
)

func check(configFilePath string) Result {
	config := parseConfig(configFilePath)
	result := Result{}

	masterTranslations := getTranslations(&result, config.Master)

	for _, locale := range config.Locales {
		for k, v := range getTranslations(&result, locale) {
			if masterValue, ok := masterTranslations[k]; ok {
				if differentInterpolations(masterValue, v) {
					result.Errors = append(result.Errors, Issue{
						Locale:  locale.Name,
						Message: fmt.Sprintf("Inconsistent interpolation for %s", k),
					})
				}
			}
		}
	}

	return result
}

func getTranslations(result *Result, l locale) translations {
	fileContents, err := translationsPerFileFor(l)
	if err != nil {
		result.Errors = append(result.Errors, Issue{
			Locale:  l.Name,
			Message: err.Error(),
		})
	}

	normalized := translations{}

	for i := 0; i < len(fileContents); i++ {
		normalized.addMap("", fileContents[i])
	}
	return normalized
}

func translationsPerFileFor(l locale) ([]map[interface{}]interface{}, error) {
	translations := make([]map[interface{}]interface{}, len(l.Files))
	for _, file := range l.Files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var i map[interface{}]interface{}
		err = yaml.Unmarshal([]byte(data), &i)
		if err != nil {
			return nil, err
		}
		translations = append(translations, i)
	}
	return translations, nil
}

func differentInterpolations(m, v string) bool {
	mi := interpolations(m)
	mv := interpolations(v)
	if len(mi) != len(mv) {
		return true
	}

	if len(mi) > 0 {
		for i, v := range mi {
			if v != mv[i] {
				return true
			}
		}
	}
	return false
}

func interpolations(text string) []string {
	re := regexp.MustCompile("%{(\\w+)}")
	interpolations := make([]string, 0)
	for _, key := range re.FindAllStringSubmatch(text, -1) {
		interpolations = append(interpolations, key[1])
	}
	sort.Strings(interpolations)
	return interpolations
}
