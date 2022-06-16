package resources

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewFile(t *testing.T) {
	tt := [] struct {
		Description string
		Want *File
	}{
		{"Getting new file", &File{}},
	}

	for _, test := range tt {
		t.Run(test.Description, func(t *testing.T) {
			got := NewFile()
			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("got: %#v, wanted: %#v", got, test.Want)
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	tt := []struct {
		Description string
		Want 		string
	} {
		{"Write testing", "testing"},
	}

	for _, test := range tt {
		file := NewFile()
		file.Content = test.Want
		buf := bytes.Buffer{}

		t.Run(test.Description, func(t *testing.T) {
			if err := file.Write(&buf); err != nil {
				t.Fatal("Error writing to buffer")
			}
			if buf.String() != test.Want {
				t.Errorf("got: %s, wanted: %s", buf.String(), test.Want)
			}
		})
	}
}