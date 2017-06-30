package interpol

import (
	"fmt"
	"strings"
)

type translations map[string]string

func (t translations) addMap(b string, content map[interface{}]interface{}) {
	for key, value := range content {
		if k, ok := key.(string); ok {
			path := fmt.Sprintf("%s.%s", b, k)
			switch v := value.(type) {
			case string:
				t.addString(path, v)
			case []interface{}:
				t.addList(path, v)
			case map[interface{}]interface{}:
				t.addMap(path, v)
			default:
				fmt.Printf("failed to parse %T - %+v at %s\n", value, value, b)
			}
		}
	}
}

func (t translations) addList(b string, lv []interface{}) {
	for i, value := range lv {
		path := fmt.Sprintf("%s.%d", b, i)
		switch v := value.(type) {
		case string:
			t.addString(path, v)
		case []interface{}:
			t.addList(path, v)
		case map[interface{}]interface{}:
			t.addMap(path, v)
		default:
			fmt.Printf("failed to parse %T - %+v at %s\n", value, value, b)
		}
	}
}

func (t translations) addString(b, s string) {
	t[strings.Join(strings.Split(b, ".")[2:], ".")] = s
}
