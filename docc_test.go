package docc

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var want = []string{
	"Title",
	"Subtitle",
	"Here is a first row.",
	"Here is a second row.",
}

func TestDecode(t *testing.T) {
	fp := "./testdata/test.docx"
	got, err := Decode(fp)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

	fp = "tmp.txt"
	_, err = Decode(fp)
	if !errors.Is(err, ErrNotSupportFormat) {
		t.Errorf("unexpected type of error: %s", err)
	}
}

func TestExtractXML(t *testing.T) {
	_, err := extractXML(filepath.Clean("./testdata/test.docx"))
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeXML(t *testing.T) {
	f, err := os.Open(filepath.Clean("./testdata/word/document.xml"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	got, err := decodeXML(f)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
