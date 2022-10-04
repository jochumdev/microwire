package client

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestMerge(t *testing.T) {
	oldConfig := NewConfig()
	newConfig := NewConfig()

	oldConfig.Plugin = "grpc"
	oldConfig.ContentType = "application/protobuf"
	oldConfig.Merge(newConfig)

	expect(t, oldConfig.ContentType, "application/protobuf")
}
