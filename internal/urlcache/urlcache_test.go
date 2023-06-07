package urlcache

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	want := []string{
		"http://google.com",
		"http://youtube.com",
		"http://facebook.com",
		"http://baidu.com",
		"http://wikipedia.org",
		"http://qq.com",
		"http://taobao.com",
	}
	got, err := ReadFile("list_test.txt")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	for i, v := range got {
		if want[i] != v {
			t.Errorf("wanted: %v, got: %v", want[i], v)
		}
	}
}
