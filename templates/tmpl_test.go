package templates

import (
	"testing"
)

func TestTmpl(t *testing.T) {
	tcs := []struct {
		file string
		size int64
	}{
		{"bootstrap/bootstrap.min.css", 153402},
		{"default/home.html", 2089},
	}
	for _, tc := range tcs {
		fss, err := FS.Open(tc.file)
		if err != nil {
			t.Error(err)
		}
		fi, err := fss.Stat()
		if err != nil {
			t.Error(err)
		}
		if tc.size != fi.Size() {
			t.Errorf("want: %d\n got: %d", tc.size, fi.Size())
		}
	}
}
