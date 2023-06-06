package urlcache

import (
	"bufio"
	"net/url"
	"os"
	"log"
)



func ReadFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()
	c := []string{}
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		u, err := CleanURL(fs.Text())
		if err != nil {
			log.Printf("got an error reading some line: %v", err)
		}
		c = append(c, u)
	}

	return c, nil
}

func CleanURL(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	if u.Scheme != "https" {
		u.Scheme = "http"
	}
	return u.String(), nil
}
