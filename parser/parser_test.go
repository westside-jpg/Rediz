package parser_test

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"github.com/westside-jpg/Rediz/parser"
	"github.com/westside-jpg/Rediz/types"
)

func WrapToReader(input string) bufio.Reader {
	return *bufio.NewReader(strings.NewReader(input))
}

func TestRequestParser(t *testing.T) {
	tests := []struct {
		name    string
		input   bufio.Reader
		want    types.Request
		wantErr bool
	}{
		{
			name:  "SET parsing with 2 elements",
			input: WrapToReader("*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$1\r\n5\r\n"),
			want: types.Request{
				Command: "SET",
				Data:    []string{"hello", "5"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.RequestParser(&tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Error %v, expected error %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Receive %v, wait %v", got, tt.want)
			}
		})
	}
}
