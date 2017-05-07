package hp4g

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestReadWrite(t *testing.T) {
	testReadWrite(t, "Vanishment this world")
	testReadWrite(t, "")
	testReadWrite(t, strings.Repeat("a", 1024))
}

func testReadWrite(t *testing.T, text string) {
	buf := new(bytes.Buffer)

	err := Write(buf, []byte(text))
	if err != nil {
		t.Error(err, "for", text)
	}

	got, err := Read(buf)
	if err != nil {
		if err == io.EOF && text == "" {
			if string(got) != "" {
				t.Error("read", got, "on EOF")
			}
		} else {
			t.Error(err, "for", text)
		}
	}

	if string(got) != text {
		t.Error("get", got, "want", text)
	}
}
