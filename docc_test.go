package docc

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

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
	want := []string{
		"Title",
		"Subtitle",
		"Here is a first row.",
		"Here is a second row.",
	}
	got, err := decodeXML(f)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
