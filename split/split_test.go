package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	var got = Split("a:b:c", ":")
	var want = []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
