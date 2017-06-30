package interpol

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	expectedConfig := config{
		Master: locale{
			Name: "en",
			Files: []string{
				"fixtures/empty-master.yml",
			},
		},
		Locales: []locale{
			locale{
				Name: "pt",
				Files: []string{
					"fixtures/empty-master.yml",
				},
			},
		},
	}

	c := parseConfig("fixtures/.empty-master.yml")
	if !reflect.DeepEqual(c, expectedConfig) {
		t.Fatal("Fail")
	}
}
