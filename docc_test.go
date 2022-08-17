package docc

import (
	"path/filepath"
	"reflect"
	"testing"
)

var expectedContent = map[string][]string{
	"./testdata/test.docx":               {"", "Title Subtitle Here is a first row. Here is a second row. ", ""},
	"./testdata/test_header_footer.docx": {"test header ", "Title Subtitle Here is a first row. Here is a second row. ", "test footer "},
}

func TestReadAll(t *testing.T) {
	for fileName, expectContent := range expectedContent {
		fp := filepath.Clean(fileName)
		r, err := NewReader(fp)
		if err != nil {
			panic(err)
		}
		defer r.Close()
		header, content, footer, err := r.ReadAllFiles()
		if err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(expectContent[0], header) {
			t.Errorf("want %v, got %v for fileName %v", expectContent[0], header, fileName)
		}
		if !reflect.DeepEqual(expectContent[1], content) {
			t.Errorf("want %v, got %v for fileName %v", expectContent[1], content, fileName)
		}
		if !reflect.DeepEqual(expectContent[2], footer) {
			t.Errorf("want %v, got %v for fileName %v", expectContent[2], footer, fileName)
		}
	}
}
