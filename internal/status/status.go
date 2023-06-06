package status

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type UrlData struct {
    url string
    t time.Duration
}

// getTime получает время до выполнения всей транзакции
func (ud *UrlData) getTime() {
    c, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    req, _ := http.NewRequestWithContext(c, "GET", ud.url, nil)

    start := time.Now()
    if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
				ud.t = -1
                return
    }
    ud.t = time.Since(start)
}
// getTime получает время до выполнения всех транзакций в списке
func GetTime(urll []string) (map[string]time.Duration) {
    m := make(map[string]time.Duration)
    var mx sync.Mutex
    var wg sync.WaitGroup
    for _, v := range(urll) {
        wg.Add(1)
        go func(url string) {
            ud := &UrlData{
                url: url,
            }
            ud.getTime()
            mx.Lock()
            m[ud.url] = ud.t
            mx.Unlock()
            wg.Done()
        } (v)
    }
    wg.Wait()
    return m
}

func SearchTime(m map[string]time.Duration, url string) string {
    v, ok := m[url]
    if ok {
        return fmt.Sprint(v.Milliseconds())
    } else {
        return fmt.Sprint("Couldn't reach the website")
    }
}

func MaxTime(m map[string]time.Duration) string {
    var max time.Duration = -1
    for _, v := range(m) {
        if v > max && v > 0 { // Намеренно не рассматривается состояние где нет соединения пока есть хоть одно нормальное соединение
            max = v
        }
    }
    return fmt.Sprint(max.Milliseconds())
}

func MinTime(m map[string]time.Duration) string {
    var min time.Duration = time.Minute
    for _, v := range(m) {
        if v < min && v > 0 {
            min = v
        }
    }
    return fmt.Sprint(min.Milliseconds())
}
