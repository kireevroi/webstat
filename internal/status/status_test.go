package status

import (
	"testing"
)

func TestGetTime(t *testing.T) {
	u := &UrlData{
		url: "http://google.com",
	}
	u.getTime()
	if u.t == -1 {
		t.Errorf("wanted no error, got one")
	}
}

func TestGetTimeError(t *testing.T) {
	u := &UrlData{
		url: "http://exampledoesntexistreallyy.io",
	}
	u.getTime()
	if u.t != -1{
		t.Errorf("wanted an error, got nothing")
	}
}

func TestGetTimeArray(t *testing.T) {
	c := []string{
		"http://google.com",
		"http://youtube.com",
		"http://facebook.com",
		"http://baidu.com",
		"http://wikipedia.org",
		"http://qq.com",
		"http://taobao.com",
	}
	got := GetTime(c)
	for k, v := range(got) {
		t.Logf("time to reach %v equals %v", k, v)
	}
}



