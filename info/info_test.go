package info

import (
	"reflect"
	"reptile/handle"
	"testing"
)

func TestParseContent(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want handle.ParseRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseContent(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
