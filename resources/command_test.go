package resources

import (
	"bytes"
	"testing"
)

func TestNewCommand(t *testing.T) {
	tt := []struct {
		name     string
		command  string
		args     []string
		expected string
	}{
		{name: "echo", command: "echo", args: []string{"testing", "hi"}, expected: "testing hi\n"},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			var b bytes.Buffer
			if err := NewCommand(v.command).SetArgs(v.args...).SetOutput(&b).Apply(); err != nil {
				t.Errorf("error running command: %v", err)
			}

			if !bytes.Equal(b.Bytes(), []byte(v.expected)) {
				t.Errorf("\nexpected %b \nbut got %b", []byte(v.expected), b.Bytes())
			}

		})
	}

}
