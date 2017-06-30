package interpol

import (
	"fmt"
	"os"
	"strings"
)

type translations map[string]string

func (t translations) addMap(b string, content map[interface{}]interface{}) {
	for key, value := range content {
		if k, ok := key.(string); ok {
			path := fmt.Sprintf("%s.%s", b, k)
			switch v := value.(type) {
			case []interface{}:
				t.addList(path, v)
			case map[interface{}]interface{}:
				t.addMap(path, v)
			default:
				t.addOther(path, v)
			}
		}
	}
}

func (t translations) addList(b string, lv []interface{}) {
	for i, value := range lv {
		path := fmt.Sprintf("%s.%d", b, i)
		switch v := value.(type) {
		case []interface{}:
			t.addList(path, v)
		case map[interface{}]interface{}:
			t.addMap(path, v)
		default:
			t.addOther(path, v)
		}
	}
}

func (t translations) addOther(b string, v interface{}) {
	if s, ok := v.(string); ok {
		t[strings.Join(strings.Split(b, ".")[2:], ".")] = s
	} else if os.Getenv("DEBUG") != "" {
		fmt.Printf("failed to parse %T - %+v at %s\n", v, v, b)
	}
}
