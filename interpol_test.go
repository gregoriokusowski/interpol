package interpol

import (
	"reflect"
	"testing"
)

func TestCheckWithHappyPath(t *testing.T) {
	result := check("fixtures/.happy-path")

	if len(result.Errors) != 0 {
		t.Fatal("Should have no errors")
	}
}

func TestCheckWithWrongInterpolations(t *testing.T) {
	result := check("fixtures/.wrong-interpolation")

	if len(result.Errors) == 0 {
		t.Fatal("Should have errors")
	}
}

func TestIntepolations(t *testing.T) {
	if len(interpolations("")) != 0 {
		t.Fatal("should be empty")
	}

	if len(interpolations("some string")) != 0 {
		t.Fatal("should be empty")
	}

	expected := []string{"another", "thing"}
	translation := "some %{thing} and %{another}... but not ${everything}"
	if !reflect.DeepEqual(interpolations(translation), expected) {
		t.Fatal("expected to have thing and another")
	}
}
